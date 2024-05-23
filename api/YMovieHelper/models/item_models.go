package models

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/utils"
)

type ItemModel struct {
	itemRepo IItemRepository
}

type IItemRepository interface {
	GetSingleItems(projectID int) ([]SingleItemForSelect, error)
	GetMultipleItems(projectID int) ([]MultipleItemForSelect, error)
	GetDynamicItems(projectID int) ([]DynamicItemForSelect, error)
	CreateSingleItem(SingleItemForInsert) error
	CreateMultipleItem(MultipleItemForInsert) error
	CreateDynamicItem(projectID int, itemType string, fileUrl string, itemPathInPC string, name string) error
	UpdateSingleItem(data *SingleItemForUpdate) error
	UpdateMultipleItem(data *MultipleItemForUpdate) error
	UpdateDynamicItem(data *DynamicItemForUpdate) error
	GetSingleItemType(typeOnYMMP string) (string, error)
	DeleteSingleItem(itemID int) error
	DeleteMultipleItem(itemID int) error
	DeleteDynamicItem(itemID int) error

	GetFileUrlForSingleItem(itemID int) (string, error)
	GetFileUrlForMultipleItem(itemID int) (string, error)
	GetFileUrlForDynamicItem(itemID int) (string, error)
}

type SingleItemForInsert struct {
	T200ID    int    `db:"t200_id"`
	M301_Name string `db:"m301_name"`
	ItemPath  string `db:"item_path"`
	Name      string `db:"name"`
	Length    int    `db:"length"`
}

type MultipleItemForInsert struct {
	T200ID       int    `db:"t200_id"`
	ItemPath     string `db:"item_path"`
	Name         string `db:"name"`
	CountOfItems int    `db:"count_of_items"`
}

type SingleItemForSelect struct {
	ID       int    `db:"id" json:"id"`
	ItemType string `db:"m301_name" json:"itemType"`
	ItemPath string `db:"item_path" json:"-"`
	Name     string `db:"name" json:"name"`
	Length   int    `db:"length" json:"length"`
}

type MultipleItemForSelect struct {
	ID           int    `db:"id" json:"id"`
	ItemPath     string `db:"item_path" json:"-"`
	Name         string `db:"name" json:"name"`
	CountOfItems int    `db:"count_of_items" json:"countOfItems"`
}
type DynamicItemForSelect struct {
	ID           int    `db:"id" json:"id"`
	ItemType     string `db:"m301_name" json:"itemType"`
	ItemUrl      string `db:"item_url" json:"-"`
	ItemPathInPC string `db:"item_path_in_pc" json:"-"`
	Name         string `db:"name" json:"name"`
}
type SingleItemForUpdate struct {
	ID     int     `json:"id"`
	Length *int    `json:"length"`
	Name   *string `json:"name"`
}

type MultipleItemForUpdate struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

type DynamicItemForUpdate struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

type ItemDescriptions struct {
	ID           string
	ItemType     string
	Descriptions []string
	Length       int
}

func NewItemModel(itemRepo IItemRepository) ItemModel {
	model := ItemModel{itemRepo: itemRepo}
	return model
}

func (model ItemModel) GetSingleItemDescriptionsAndSave(ctx *gin.Context, data *domains.YMMP) ([]ItemDescriptions, error) {
	var errorMessages []error
	items := data.Timeline.Items
	var descriptions []ItemDescriptions

	// そもそもアイテムが含まれていなければエラー
	if len(data.Timeline.Items) == 0 {
		customError := &utils.CustomError{
			FrontMsg: "ymmpファイルの中にアイテムが含まれていません！最低でも1つはアイテムを配置してください！",
			BackMsg:  "No items.",
		}
		return nil, fmt.Errorf("ItemModel.GetSingleItemDescriptionsAndSave: %w", customError)
	}

	// アイテムにVoiceCacheが含まれていればエラー
	for _, item := range data.Timeline.Items {
		if item.VoiceCache != nil {
			customError := &utils.CustomError{
				FrontMsg: "ymmpファイルにボイスキャッシュが含まれています！（詳しくはHow to ページをご覧ください。）",
				BackMsg:  "Include Voice cache.",
			}
			return nil, fmt.Errorf("ItemModel.GetSingleItemDescriptionsAndSave: %w", customError)
		}
	}

	for _, item := range items {
		uuid := uuid.New().String()
		description, err := model.makeItemDescriptions(item, uuid)

		if err != nil {
			errorMessages = append(errorMessages, fmt.Errorf("ItemModel.GetSingleItemDescriptions: %w\n", err))
			continue
		}

		filePath := fmt.Sprintf("./temp/single_items/%s.json", uuid)
		file, err := json.Marshal(item)
		if err != nil {
			errorMessages = append(errorMessages, fmt.Errorf("ItemModel.GetSingleItemDescriptions: failed to json.Marshal:%w\n", err))
			continue
		}
		err = utils.SaveFile(bytes.NewReader(file), filePath)
		if err != nil {
			errorMessages = append(errorMessages, fmt.Errorf("ItemModel.GetSingleItemDescriptions: failed to json.Marshal:%w\n", err))
			continue
		}
		descriptions = append(descriptions, description)
	}
	//エラーを一つの文字列として結合
	if len(errorMessages) > 0 {
		errorMessageStrings := make([]string, len(errorMessages))
		for i, err := range errorMessages {
			errorMessageStrings[i] = err.Error()
		}
		return descriptions, fmt.Errorf("ItemModel. GetSingleItemDescriptionsAndSave: multiple errors: \n%s", strings.Join(errorMessageStrings, "; "))
	}

	return descriptions, nil
}

func (model ItemModel) UploadSingleItem(ctx context.Context, projectID int, name string, id string, itemType string, length int) error {
	oldFilePath := fmt.Sprintf("./temp/single_items/%s.json", id)
	newFilePath := fmt.Sprintf("./data/single_items/%s.json", id)

	err := utils.MoveFile(oldFilePath, newFilePath)
	if err != nil {
		return fmt.Errorf("ItemModel.UploadSingleItem: %w", err)
	}

	singleItem := SingleItemForInsert{T200ID: projectID, M301_Name: itemType, ItemPath: newFilePath, Name: name, Length: length}

	err = model.itemRepo.CreateSingleItem(singleItem)
	if err != nil {
		return fmt.Errorf("ItemModel.UploadSingleItem: %w", err)
	}

	return nil
}

// FIXED: わざわざ構造体を定義する必要性が薄いので、引数を全て取るようにしている、勉強後にどちらの方がベストプラクティスか判断
func (model ItemModel) UploadDynamicItem(ctx *gin.Context, data *domains.YMMP, projectID int, itemType string, name string) error {
	items := data.Timeline.Items
	if len(items) == 0 || len(items) > 1 {
		customError := &utils.CustomError{
			FrontMsg: "ymmpファイルの中にアイテムはただ1つだけ含めてください！",
			BackMsg:  "The number of items must be one! ",
		}
		return fmt.Errorf("ItemModel.UploadDynamicItem: %w", customError)
	}
	//ファイルの親ディレクトリの親ディレクトリを抽出
	components := strings.Split(items[0].FilePath, "\\")
	if len(components) < 3 {
		customError := &utils.CustomError{
			FrontMsg: "ymmpファイル内のアイテムのファイルパスがおかしいようです。「Software/Project/DynamicItem/[Dynamic Item]/素材名」のフォルダ構成を守ってください。",
			BackMsg:  "Not enough components in path.",
		}
		return fmt.Errorf("ItemModel.UploadDynamicItem: %w", customError)
	}
	itemPathInPC, err := utils.Encrypt(strings.Join(components[:len(components)-3], "\\"))
	if err != nil {
		return fmt.Errorf("ItemModel.UploadDynamicItem: %v", err)
	}

	ymmpItemType, err := model.itemRepo.GetSingleItemType(*items[0].Type)
	if err != nil {
		return fmt.Errorf("ItemModel.UploadDynamicItem: %v", err)
	}

	if itemType != ymmpItemType {
		customError := &utils.CustomError{
			FrontMsg: "宣言されたアイテムのカテゴリと、ファイル内に含まれるアイテムのカテゴリが異なるようです。",
			BackMsg:  "item type in ymmp != item type fron client.",
		}
		return fmt.Errorf("ItemModel.UploadDynamicItem: %w", customError)
	}

	uploadFile, err := json.Marshal(items[0])
	if err != nil {
		customError := &utils.CustomError{
			FrontMsg: "不正なymmpファイルです。心当たりがない場合は問い合わせフォームまでご連絡をお願いいたします。",
			BackMsg:  "failed to json.Marshal",
		}
		return fmt.Errorf("ItemModel.UploadDynamicItem: %w", customError)
	}
	//UUIDで一意のファイルURLを作成
	filePath := fmt.Sprintf("./data/dynamic_items/%s.json", uuid.New())
	err = utils.SaveFile(bytes.NewReader(uploadFile), filePath)
	if err != nil {
		return fmt.Errorf("ItemModel.UploadDynamicItem: failed to save file:%w", err)
	}

	err = model.itemRepo.CreateDynamicItem(projectID, itemType, filePath, itemPathInPC, name)
	if err != nil {
		return fmt.Errorf("ItemModel.UploadDynamicItem: failed to CreateDynamicItem:%w", err)
	}

	return nil

}

func (model ItemModel) UploadMultipleItem(ctx *gin.Context, data *domains.YMMP, projectID int, name string) error {
	items := data.Timeline.Items

	// そもそもアイテムが含まれていなければエラー
	if len(items) == 0 {
		customError := &utils.CustomError{
			FrontMsg: "ymmpファイルの中にアイテムが含まれていません！",
			BackMsg:  "No items.",
		}
		return fmt.Errorf("ItemModel.UploadMultipleItem: %w", customError)
	}

	// アイテムにVoiceCacheが含まれていればエラー
	for _, item := range data.Timeline.Items {
		if item.VoiceCache != nil {
			customError := &utils.CustomError{
				FrontMsg: "ymmpファイルの中にボイスキャッシュが含まれているようです！（詳しくはHow toページをご覧ください。）",
				BackMsg:  "Include Voice Cache.",
			}
			return fmt.Errorf("ItemModel.UploadMultipleItem: %w", customError)
		}
	}

	countOfItems := len(items)
	resetMultipleItemFrameAndLayer(items)
	uploadFile, err := json.Marshal(items)
	if err != nil {
		customError := &utils.CustomError{
			FrontMsg: "不正なymmpファイルです。心当たりがない場合は問い合わせフォームまでご連絡をお願いいたします。",
			BackMsg:  "failed to json.Marshal",
		}
		return fmt.Errorf("ItemModel.UploadMultipleItem: %w", customError)
	}
	//UUIDで一意のファイルURLを作成
	filePath := fmt.Sprintf("./data/multiple_items/%s.json", uuid.New())
	err = utils.SaveFile(bytes.NewReader(uploadFile), filePath)
	if err != nil {
		return fmt.Errorf("ItemModel.UploadMultipleItem: failed to save file:%w", err)
	}

	multipleItem := MultipleItemForInsert{T200ID: projectID, ItemPath: filePath, Name: name, CountOfItems: countOfItems}
	err = model.itemRepo.CreateMultipleItem(multipleItem)
	if err != nil {
		return fmt.Errorf("ItemModel.UploadMultipleItem: failed to CreateMultipleItem:%w", err)
	}

	return nil

}

func (model ItemModel) GetItems(projectID int) ([]SingleItemForSelect, []MultipleItemForSelect, []DynamicItemForSelect, error) {
	singleItems, err := model.itemRepo.GetSingleItems(projectID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("ItemModel.GetItems: %w", err)
	}

	multipleItems, err := model.itemRepo.GetMultipleItems(projectID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("ItemModel.GetItems: %w", err)
	}

	dynamicItems, err := model.itemRepo.GetDynamicItems(projectID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("ItemModel.GetItems: %w", err)
	}

	return singleItems, multipleItems, dynamicItems, nil
}

func (model ItemModel) UpdateSingleItem(data *SingleItemForUpdate) error {
	err := model.itemRepo.UpdateSingleItem(data)
	if err != nil {
		return fmt.Errorf("ItemModel.UpdateSingleItem: %w", err)
	}

	return nil
}
func (model ItemModel) UpdateMultipleItem(data *MultipleItemForUpdate) error {
	err := model.itemRepo.UpdateMultipleItem(data)
	if err != nil {
		return fmt.Errorf("ItemModel.UpdateMultipleItem: %w", err)
	}

	return nil
}

func (model ItemModel) UpdateDynamicItem(data *DynamicItemForUpdate) error {
	err := model.itemRepo.UpdateDynamicItem(data)
	if err != nil {
		return fmt.Errorf("ItemModel.UpdateDynamicItem: %w", err)
	}

	return nil
}

func (model ItemModel) DeleteSingleItem(ctx context.Context, itemID int) error {
	fileUrl, err := model.itemRepo.GetFileUrlForSingleItem(itemID)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteSingleItem: %w", err)
	}
	err = model.itemRepo.DeleteSingleItem(itemID)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteSingleItem:%w", err)
	}

	err = utils.DeleteFile(fileUrl)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteSingleItem: %w", err)
	}

	return err
}
func (model ItemModel) DeleteMultipleItem(ctx context.Context, itemID int) error {
	fileUrl, err := model.itemRepo.GetFileUrlForMultipleItem(itemID)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteMultipleItem: %w", err)
	}
	err = model.itemRepo.DeleteMultipleItem(itemID)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteMultipleItem:%w", err)
	}

	err = utils.DeleteFile(fileUrl)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteMultipleItem: %w", err)
	}

	return err
}
func (model ItemModel) DeleteDynamicItem(ctx context.Context, itemID int) error {
	fileUrl, err := model.itemRepo.GetFileUrlForDynamicItem(itemID)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteDynamicItem: %w", err)
	}
	err = model.itemRepo.DeleteDynamicItem(itemID)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteDynamicItem:%w", err)
	}

	err = utils.DeleteFile(fileUrl)
	if err != nil {
		return fmt.Errorf("ItemModel.DeleteDynamicItem: %w", err)
	}

	return err
}

// 送られてきたアイテムの属性からItemDescriptionsを作成する関数
func (model ItemModel) makeItemDescriptions(item *domains.ItemOnYMMP, id string) (ItemDescriptions, error) {
	itemType, err := model.itemRepo.GetSingleItemType(*item.Type)
	if err != nil {
		return ItemDescriptions{}, fmt.Errorf("ItemModel.MakeItemDescriptions: %w", err)
	}

	length := item.Length
	switch itemType {
	case "ボイスアイテム":
		characterName := "キャラクター名が設定されていないようです。"
		if item.CharacterName != nil {
			characterName = *item.CharacterName
		}
		serif := "(空のセリフ)"
		if item.Serif != nil {
			serif = *item.Serif
		}
		return ItemDescriptions{
			ItemType:     itemType,
			Descriptions: []string{characterName, serif},
			ID:           id,
			Length:       length,
		}, nil

	case "テキストアイテム":
		text := "(空のテキストアイテム)"
		if item.Text != nil {
			text = *item.Text
		}
		return ItemDescriptions{
			ItemType:     itemType,
			Descriptions: []string{text},
			ID:           id,
			Length:       length,
		}, nil

	case "ビデオアイテム", "オーディオアイテム", "画像アイテム":
		filePath := item.FilePath
		paths := strings.Split(filePath, "\\")
		fileName := paths[len(paths)-1]
		return ItemDescriptions{
			ItemType:     itemType,
			Descriptions: []string{fileName},
			ID:           id,
			Length:       length,
		}, nil

	case "立ち絵アイテム", "表情アイテム":
		characterName := "キャラクター名が設定されていないようです。"
		if item.CharacterName != nil {
			characterName = *item.CharacterName
		}
		return ItemDescriptions{
			ItemType:     itemType,
			Descriptions: []string{characterName},
			ID:           id,
			Length:       length,
		}, nil

	case "図形アイテム", "エフェクトアイテム", "画面の複製", "グループ制御":
		return ItemDescriptions{
			ItemType:     itemType,
			Descriptions: []string{},
			ID:           id,
			Length:       length,
		}, nil

	}
	return ItemDescriptions{}, errors.New("ItemModel.MakeItemDescriptions: unknown item type")
}

// 全てのitemのframeを0基準にする
func resetMultipleItemFrameAndLayer(items []*domains.ItemOnYMMP) {
	if len(items) == 0 {
		return
	}

	// 最小のフレーム値を見つける
	minFrame := items[0].Frame
	for _, item := range items {
		if item.Frame < minFrame {
			minFrame = item.Frame
		}
	}

	// 各フレームから最小のフレーム値を減算する
	for i := range items {
		(items)[i].Frame -= minFrame
	}

	minLayer := (items)[0].Layer
	for _, item := range items {
		if item.Layer < minLayer {
			minLayer = item.Layer
		}
	}

	for i := range items {
		(items)[i].Layer -= minLayer
	}

}
