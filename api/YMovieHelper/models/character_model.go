package models

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/utils"
)

type ICharacterRepository interface {
	Create(softwareID int, name string) error
	GetCharacters(softwareID int) ([]CharacterForSelect, error)
	Update(data *CharacterForUpdate) error
	Delete(characterID int) error

	CreateEmotion(characterID int, emotionName string, fileUrl string) error
	GetEmotions(characterID int) (characterName string, emotions []EmotionForSelect, err error)
	UpdateEmotion(data *EmotionForUpdate) error
	DeleteEmotion(emotionID int) error

	GetFileUrls(characterID int) ([]string, error)

	GetFileUrlForEmotion(emotionID int) (string, error)
}
type CharacterModel struct {
	charaRepo ICharacterRepository
}

type CharacterForSelect struct {
	ID       int    `db:"id"`
	IsEmpty  bool   `db:"is_empty"`
	Name     string `db:"name"`
	Emotions []EmotionForSelect
}

type CharacterForUpdate struct {
	ID      int     `json:"id"`
	Name    *string `json:"name"`
	IsEmpty *bool   `json:"isEmpty"`
}

type EmotionForSelect struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type EmotionForUpdate struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

func NewCharacterModel(charaRepo ICharacterRepository) CharacterModel {
	model := CharacterModel{charaRepo: charaRepo}
	return model
}

func (model CharacterModel) CreateCharacters(ymmpData *domains.YMMP, softwareID int) error {
	var errorMessages []error
	characters := ymmpData.Characters
	// アイテムにVoiceCacheが含まれていればエラー
	for _, item := range ymmpData.Timeline.Items {
		if item.VoiceCache != nil {
			customError := &utils.CustomError{
				FrontMsg: "ymmpファイルの中にボイスキャッシュが含まれているようです。（詳しくはHow toページをご覧ください。）",
				BackMsg:  "Include Voice cache.",
			}
			return fmt.Errorf("CharacterModel.CreateCharacters: %w", customError)
		}
	}

	// CharacterNameが空文字列のものがあればエラー
	for _, character := range characters {
		if character.Name == "" {
			errorMessages = append(errorMessages, fmt.Errorf("CharacterModel.CreateCharacters: Empty Character name.\n"))
			continue
		}
		err := model.charaRepo.Create(softwareID, character.Name)
		if err != nil {
			errorMessages = append(errorMessages, fmt.Errorf("CharacterModel.CreateCharacters: %w\n", err))
		}
	}

	if len(errorMessages) > 0 {
		errorMessageStrings := make([]string, len(errorMessages))
		for i, err := range errorMessages {
			errorMessageStrings[i] = err.Error()
		}
		customError := &utils.CustomError{
			FrontMsg: "名前が空白のキャラが含まれているようです。最低でも1文字は入力してください。（サーバーエラーの可能性があります。心当たりがない場合は時間をおいてから試してください。）",
			BackMsg:  fmt.Sprintf("multiple errors: \n%s", strings.Join(errorMessageStrings, "; ")),
		}
		return fmt.Errorf("CharacterModel.CreateCharacters: %w", customError)

	}

	return nil
}

func (model CharacterModel) GetCharacters(softwareID int) ([]CharacterForSelect, error) {
	characters, err := model.charaRepo.GetCharacters(softwareID)
	if err != nil {
		return characters, fmt.Errorf("CharacterModel.GetCharacter: %v", err)
	}
	return characters, nil
}

func (model CharacterModel) UpdateCharacter(data *CharacterForUpdate) error {
	err := model.charaRepo.Update(data)
	if err != nil {
		return fmt.Errorf("CharacterModel.UpdateCharacter: %v", err)
	}

	return nil
}

func (model CharacterModel) DeleteCharacter(ctx context.Context, characterID int) error {
	var errorMessages []error
	fileUrls, err := model.charaRepo.GetFileUrls(characterID)
	if err != nil {
		return fmt.Errorf("CharacterModel.Deletecharacter: %w", err)
	}

	err = model.charaRepo.Delete(characterID)
	if err != nil {
		return fmt.Errorf("CharacterModel.DeleteCharacter: %v", err)
	}

	for _, fileUrl := range fileUrls {
		err := utils.DeleteFile(fileUrl)
		if err != nil {
			errorMessages = append(errorMessages, err)
		}
	}

	if len(errorMessages) > 0 {
		var errorMessage strings.Builder
		for _, err := range errorMessages {
			errorMessage.WriteString(fmt.Sprintf("CharacterModel.DeleteCharacter: %v\n", err))
		}
		return errors.New(errorMessage.String())
	}

	return err
}

func (model CharacterModel) CreateEmotions(ctx context.Context, data *domains.YMMP, characterID int) error {
	// 空セリフ、あるいはボイスアイテムではないもの、あるいはVoiceCacheが混じっていないかのバリデーション
	for _, item := range data.Timeline.Items {
		if *item.Type != "YukkuriMovieMaker.Project.Items.VoiceItem, YukkuriMovieMaker" {
			customError := &utils.CustomError{
				FrontMsg: "ymmpファイルの中にボイスアイテム以外が含まれているようです。再度ご確認ください。",
				BackMsg:  "Include Not VoiceItem",
			}
			return fmt.Errorf("models.CharacterModel.CreateEmotions: %w", customError)
		}

		if utils.IsEmptyOrWhitespace(*item.Serif) {
			customError := &utils.CustomError{
				FrontMsg: "セリフが空のボイスアイテムが含まれています。最低でも空白以外の1つ以上の文字を入力してください。",
				BackMsg:  "Include empty serif ",
			}
			return fmt.Errorf("models.CharacterModel.CreateEmotions: %w", customError)
		}
		// アイテムにVoiceCacheが含まれていればエラー
		if item.VoiceCache != nil {
			customError := &utils.CustomError{
				FrontMsg: "ymmpファイルの中にボイスキャッシュが含まれているようです。（詳しくはHow toページをご覧ください。）",
				BackMsg:  "Include Voice cache.",
			}
			return fmt.Errorf("models.CharacterModel.CreateEmotions: %w", customError)
		}
	}
	for _, item := range data.Timeline.Items {
		file, err := json.Marshal(item.TachieFaceParameter)
		if err != nil {
			return fmt.Errorf("CharacterModel.CreateEmotions: %w", err)
		}

		filePath := fmt.Sprintf("./data/emotions/%s.json", uuid.New().String())

		err = utils.SaveFile(bytes.NewReader(file), filePath)
		if err != nil {
			return fmt.Errorf("CharacterModel.CreateEmotions: %w", err)
		}

		err = model.charaRepo.CreateEmotion(characterID, *item.Serif, filePath)
		if err != nil {
			return fmt.Errorf("CharacterModel.CreateEmotions: %w", err)
		}
	}

	return nil
}

func (model CharacterModel) GetEmotions(characterID int) (string, []EmotionForSelect, error) {
	characterName, emotions, err := model.charaRepo.GetEmotions(characterID)
	if err != nil {
		return "", nil, fmt.Errorf("CharacterModel.GetEmotion: %w", err)
	}

	return characterName, emotions, nil
}

func (model CharacterModel) UpdateEmotion(data *EmotionForUpdate) error {
	err := model.charaRepo.UpdateEmotion(data)
	if err != nil {
		return fmt.Errorf("CharacterModel.UpdateEmotion: %v", err)
	}

	return nil
}

func (model CharacterModel) DeleteEmotion(ctx context.Context, emotionID int) error {
	fileUrl, err := model.charaRepo.GetFileUrlForEmotion(emotionID)
	if err != nil {
		return fmt.Errorf("CharacterModel.DeleteEmotion: %w", err)
	}
	err = model.charaRepo.DeleteEmotion(emotionID)
	if err != nil {
		return fmt.Errorf("CharacterModel.DeleteEmotion: %w", err)
	}

	err = utils.DeleteFile(fileUrl)
	if err != nil {
		return fmt.Errorf("CharacterModel.DeleteEmotion: %w", err)
	}

	return err
}
