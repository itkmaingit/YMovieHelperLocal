package repository

import (
	"fmt"
	"strings"

	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/utils"
)

type CharacterRepository struct {
}

func NewCharacterRepository() models.ICharacterRepository {
	return CharacterRepository{}
}

func (CharacterRepository) ExistsCharacter(softwareID int, characterID int) (bool, error) {
	db := GetDB()
	var exists bool
	selectQuery := `SELECT EXISTS(
									SELECT *
									FROM t100_software
									INNER JOIN t110_character
									ON t100_software.id = t110_character.t100_id
									WHERE t100_software.id = ? AND t110_character.id =?)`
	err := db.Get(&exists, selectQuery, softwareID, characterID)
	if err != nil {
		return exists, fmt.Errorf("CharacterRepository.ExistsCharacter: failed to check if project exists: %w", err)
	}
	return exists, err
}

func (CharacterRepository) Create(softwareID int, name string) error {
	db := GetDB()
	var count int
	checkQuery := "SELECT COUNT(*) FROM t110_character WHERE t100_id = ?"
	insertQuery := "INSERT INTO t110_character (t100_id, name) VALUES (?, ?)"

	err := db.Get(&count, checkQuery, softwareID)
	if err != nil || count >= 10 {
		return fmt.Errorf("CharacterRepository.CreateCharacter: failed to insert data: %w", err)
	}

	_, err = db.Exec(insertQuery, softwareID, name)
	if err != nil {
		return fmt.Errorf("CharacterRepository.CreateCharacter: %w", err)
	}
	return err
}

func (CharacterRepository) GetCharacters(softwareID int) ([]models.CharacterForSelect, error) {
	db := GetDB()
	var characters []models.CharacterForSelect
	selectQueryForT110 := "SELECT id, is_empty, name FROM t110_character WHERE t100_id=?"
	selectQueryForT111 := "SELECT id, name FROM t111_emotion WHERE t110_id=?"
	err := db.Select(&characters, selectQueryForT110, softwareID)
	if err != nil {
		return nil, fmt.Errorf("CharacterRepository.GetCharacters: %w", err)
	}

	for index, character := range characters {
		err := db.Select(&characters[index].Emotions, selectQueryForT111, character.ID)
		if err != nil {
			return nil, fmt.Errorf("CharacterRepository.GetCharacters: %w", err)
		}
	}
	return characters, nil
}

func (CharacterRepository) Update(data *models.CharacterForUpdate) error {
	db := GetDB()
	var updateQuery strings.Builder
	updateQuery.WriteString("UPDATE t110_character SET ")

	var args []interface{}
	if data.Name != nil {
		updateQuery.WriteString("name = ?, ")
		args = append(args, *data.Name)
	}

	if data.IsEmpty != nil {
		updateQuery.WriteString("is_empty = ?, ")
		args = append(args, *data.IsEmpty)
	}

	query := utils.ModifyQuery(updateQuery)

	query += "WHERE id=?"
	args = append(args, data.ID)

	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("repository.CharacterRepository.Update: %w", err)
	}

	return err
}

func (CharacterRepository) GetFileUrls(characterID int) ([]string, error) {
	db = GetDB()
	var fileUrls []string

	var emotionItemUrls []string

	selectQueryForT111 := `
	SELECT t111_emotion.item_path
	FROM t111_emotion
	INNER JOIN t110_character
	ON t111_emotion.t110_id = t110_character.id
	WHERE t110_character.id = ?
`
	err := db.Select(&emotionItemUrls, selectQueryForT111, characterID)
	if err != nil {
		return nil, fmt.Errorf("CharacterRepository.GetFileUrls: %w", err)
	}

	fileUrls = append(fileUrls, emotionItemUrls...)

	return fileUrls, err
}

func (CharacterRepository) Delete(characterID int) error {
	db := GetDB()
	deleteQuery := "DELETE FROM t110_character WHERE id=?"
	_, err := db.Exec(deleteQuery, characterID)
	if err != nil {
		return fmt.Errorf("CharacterRepository.Delete: %w", err)
	}
	return nil
}

func (CharacterRepository) CreateEmotion(characterID int, emotionName string, fileUrl string) error {
	db := GetDB()
	var count int
	checkQuery := "SELECT COUNT(*) FROM t111_emotion WHERE t110_id = ?"
	insertQuery := "INSERT INTO t111_emotion (t110_id, item_path, name) VALUES (?,?,?)"

	err := db.Get(&count, checkQuery, characterID)
	if err != nil || count >= 20 {
		return fmt.Errorf("CharacterRepository.CreateEmotion: failed to insert data: %w", err)
	}
	_, err = db.Exec(insertQuery, characterID, fileUrl, emotionName)
	if err != nil {
		return fmt.Errorf("CharacterRepository.CreateEmotion: %w", err)
	}

	return err
}

func (CharacterRepository) GetEmotions(characterID int) (string, []models.EmotionForSelect, error) {
	db := GetDB()
	var characterName string
	var emotions []models.EmotionForSelect
	selectQueryForCharacterName := "SELECT name FROM t110_character WHERE id=?"
	selectQueryForEmotions := "SELECT id, name FROM t111_emotion WHERE t110_id=?"
	err := db.Get(&characterName, selectQueryForCharacterName, characterID)
	if err != nil {
		return "", nil, fmt.Errorf("CharacterRepository.GetEmotions: %w", err)
	}
	err = db.Select(&emotions, selectQueryForEmotions, characterID)
	if err != nil {
		return "", nil, fmt.Errorf("CharacterRepository.GetEmotions: %w", err)
	}

	return characterName, emotions, nil
}

func (CharacterRepository) UpdateEmotion(data *models.EmotionForUpdate) error {
	db := GetDB()
	var updateQuery strings.Builder
	updateQuery.WriteString("UPDATE t111_emotion SET ")

	var args []interface{}
	if data.Name != nil {
		updateQuery.WriteString("name = ?, ")
		args = append(args, *data.Name)
	}

	query := utils.ModifyQuery(updateQuery)
	query += "WHERE id = ?"

	args = append(args, data.ID)

	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("repository.CharacterRepository.UpdateEmotion: %w", err)
	}

	return err
}

func (CharacterRepository) GetFileUrlForEmotion(emotionID int) (string, error) {
	db = GetDB()
	var fileUrl string

	selectQueryForT111 := `
	SELECT t111_emotion.item_path
	FROM t111_emotion
	WHERE t111_emotion.id = ?
`
	err := db.Get(&fileUrl, selectQueryForT111, emotionID)
	if err != nil {
		return "", fmt.Errorf("CharacterRepository.GetFileUrlForEmotion: %w", err)
	}

	return fileUrl, err
}

func (CharacterRepository) DeleteEmotion(emotionID int) error {
	db := GetDB()
	deleteQuery := "DELETE FROM t111_emotion WHERE id=?"
	_, err := db.Exec(deleteQuery, emotionID)
	if err != nil {
		return fmt.Errorf("CharacterRepository.Delete: %w", err)
	}
	return nil
}
