package repository

import (
	"fmt"
	"strings"

	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/utils"
)

type SoftwareRepository struct {
}

func NewSoftwareRepository() models.ISoftwareRepository {
	return SoftwareRepository{}
}

func (SoftwareRepository) ExistsSoftware(softwareID int) (bool, error) {
	db := GetDB()
	var exists bool
	selectQuery := `SELECT EXISTS(
									SELECT *
									FROM t100_software
									WHERE t100_software.id = ?)`
	err := db.Get(&exists, selectQuery, softwareID)
	if err != nil {
		return exists, fmt.Errorf("failed to check if software exists: %w", err)
	}
	return exists, err
}

func (SoftwareRepository) GetSoftwares() ([]models.SoftwareAndProjects, error) {
	db := GetDB()
	var softwares []models.SoftwareAndProjects
	selectQuery := `SELECT t100_software.id, t100_software.name
									FROM t100_software`
	err := db.Select(&softwares, selectQuery)
	if err != nil {
		return softwares, fmt.Errorf("SoftwareRepository.GetSoftwares: %w", err)
	}
	return softwares, err
}

func (SoftwareRepository) Create(data *models.SoftwareForInsert) (int, error) {
	var count int
	db := GetDB()
	checkQuery := "SELECT COUNT(*) FROM t100_software"
	err := db.Get(&count, checkQuery)
	if err != nil || count >= 3 {
		return 0, fmt.Errorf("repository.SoftwareRepository.Create: failed to insert data: %w", err)
	}

	insertQuery := "INSERT INTO t100_software (name) VALUES (:name)"
	result, err := db.NamedExec(insertQuery, *data)
	if err != nil {
		return 0, fmt.Errorf("repository.SoftwareRepository.Create: failed to insert data: %w", err)
	}
	softwareID_64, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("repository.SoftwareRepository.Create: failed to get last insert id: %w", err)
	}

	softwareID := int(softwareID_64)
	return softwareID, err
}

func (SoftwareRepository) GetProjects(softwareID int) ([]models.ProjectForSelect, error) {
	db := GetDB()
	var projects []models.ProjectForSelect
	selectQuery := `SELECT t200_project.id, t200_project.name
									FROM t100_software
									INNER JOIN t200_project
									ON t100_software.id =t200_project.t100_id
									WHERE t100_software.id =?`
	err := db.Select(&projects, selectQuery, softwareID)
	if err != nil {
		return projects, fmt.Errorf("repository.SoftwareRepository.GetProjects: %w", err)
	}
	return projects, err
}

func (SoftwareRepository) Update(data *models.SoftwareForUpdate) error {
	db := GetDB()
	var updateQuery strings.Builder
	updateQuery.WriteString("UPDATE t100_software SET ")

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
		return fmt.Errorf("repository.SoftwareRepository.Update: %w", err)
	}

	return err
}

func (SoftwareRepository) GetFileUrls(softwareID int) ([]string, error) {
	db = GetDB()
	var fileUrls []string

	var singleItemUrls []string
	var multipleItemUrls []string
	var dynamicItemUrls []string
	var emotionItemUrls []string

	selectQueryForT310 := `
	SELECT t310_single_item.item_path
	FROM t310_single_item
	INNER JOIN t200_project
	ON t310_single_item.t200_id = t200_project.id
	WHERE t200_project.t100_id = ?
	`
	selectQueryForT320 := `
	SELECT t320_multiple_item.item_path
	FROM t320_multiple_item
	INNER JOIN t200_project
	ON t320_multiple_item.t200_id = t200_project.id
	WHERE t200_project.t100_id = ?
	`
	selectQueryForT330 := `
	SELECT t330_dynamic_item.item_url
	FROM t330_dynamic_item
	INNER JOIN t200_project
	ON t330_dynamic_item.t200_id = t200_project.id
	WHERE t200_project.t100_id = ?
	`
	selectQueryForT111 := `
	SELECT t111_emotion.item_path
	FROM t111_emotion
	INNER JOIN t110_character
	ON t111_emotion.t110_id = t110_character.id
	INNER JOIN t100_software
	ON t110_character.t100_id = t100_software.id
	WHERE t100_software.id = ?
	`
	err := db.Select(&singleItemUrls, selectQueryForT310, softwareID)
	if err != nil {
		return nil, fmt.Errorf("SoftwareRepository.GetFileUrls: %w", err)
	}
	err = db.Select(&multipleItemUrls, selectQueryForT320, softwareID)
	if err != nil {
		return nil, fmt.Errorf("SoftwareRepository.GetFileUrls: %w", err)
	}
	err = db.Select(&dynamicItemUrls, selectQueryForT330, softwareID)
	if err != nil {
		return nil, fmt.Errorf("SoftwareRepository.GetFileUrls: %w", err)
	}
	err = db.Select(&emotionItemUrls, selectQueryForT111, softwareID)
	if err != nil {
		return nil, fmt.Errorf("SoftwareRepository.GetFileUrls: %w", err)
	}

	fileUrls = append(fileUrls, singleItemUrls...)
	fileUrls = append(fileUrls, multipleItemUrls...)
	fileUrls = append(fileUrls, dynamicItemUrls...)
	fileUrls = append(fileUrls, emotionItemUrls...)

	return fileUrls, err
}

func (SoftwareRepository) Delete(softwareID int) error {
	db := GetDB()
	deleteQuery := `DELETE FROM t100_software WHERE t100_software.id =?`
	_, err := db.Exec(deleteQuery, softwareID)
	if err != nil {
		return fmt.Errorf("repository.SoftwareRepository.Delete: %w", err)
	}

	return nil
}
