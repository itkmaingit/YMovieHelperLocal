package repository

import (
	"fmt"
	"strings"

	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/utils"
)

type ProjectRepository struct {
}

func NewProjectRepository() models.IProjectRepository {
	return ProjectRepository{}
}

func (ProjectRepository) Create(data *models.ProjectForInsert) (int, error) {
	db := GetDB()
	var projectID_64 int64
	var projectID int
	var count int

	checkQuery := "SELECT COUNT(*) FROM t200_project WHERE t100_id = ?"
	err := db.Get(&count, checkQuery, data.T100ID)
	if err != nil || count >= 5 {
		return 0, fmt.Errorf("repoistory.ProjectRepository.Create: failed to insert data: %w", err)
	}

	insertQuery := "INSERT INTO t200_project (t100_id,name) VALUES (:t100_id,:name)"
	result, err := db.NamedExec(insertQuery, *data)
	if err != nil {
		return projectID, fmt.Errorf("repoistory.ProjectRepository.Create: %w", err)
	}
	projectID_64, err = result.LastInsertId()
	if err != nil {
		return projectID, fmt.Errorf("repoistory.ProjectRepository.Create:  %w", err)
	}
	projectID = int(projectID_64)
	return projectID, err
}

func (ProjectRepository) Update(data *models.ProjectForUpdate) error {
	db := GetDB()
	var updateQuery strings.Builder
	updateQuery.WriteString("UPDATE t200_project SET ")

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
		return fmt.Errorf("repository.ProjectRepository.Update: %w", err)
	}

	return err
}

func (ProjectRepository) Delete(projectID int) error {
	db := GetDB()
	deleteQuery := "DELETE FROM t200_project WHERE id = ?"
	_, err := db.Exec(deleteQuery, projectID)
	if err != nil {
		return fmt.Errorf("repoistory.ProjectRepository.Delete: %w", err)
	}
	return nil
}

func (ProjectRepository) ExistsProject(uid string, softwareID int, projectID int) (bool, error) {
	db := GetDB()
	var exists bool
	selectQuery := `SELECT EXISTS(
									SELECT *
									FROM t100_software
									INNER JOIN t200_project
									ON t100_software.id = t200_project.t100_id
									WHERE t100_software.id = ? AND t200_project.id =?)`
	err := db.Get(&exists, selectQuery, uid, softwareID, projectID)
	if err != nil {
		return exists, fmt.Errorf("ProjectRepository.ExistsProject: failed to check if project exists: %w", err)
	}
	return exists, err
}

func (ProjectRepository) GetFileUrls(projectID int) ([]string, error) {
	db = GetDB()
	var fileUrls []string

	var singleItemUrls []string
	var multipleItemUrls []string
	var dynamicItemUrls []string

	selectQueryForT310 := `
	SELECT t310_single_item.item_path
	FROM t310_single_item
	INNER JOIN t200_project
	ON t310_single_item.t200_id = t200_project.id
	WHERE t200_project.id = ?
	`
	selectQueryForT320 := `
	SELECT t320_multiple_item.item_path
	FROM t320_multiple_item
	INNER JOIN t200_project
	ON t320_multiple_item.t200_id = t200_project.id
	WHERE t200_project.id = ?
	`
	selectQueryForT330 := `
	SELECT t330_dynamic_item.item_url
	FROM t330_dynamic_item
	INNER JOIN t200_project
	ON t330_dynamic_item.t200_id = t200_project.id
	WHERE t200_project.id = ?
	`
	err := db.Select(&singleItemUrls, selectQueryForT310, projectID)
	if err != nil {
		return nil, fmt.Errorf("ProjectRepository.GetFileUrls: %w", err)
	}
	err = db.Select(&multipleItemUrls, selectQueryForT320, projectID)
	if err != nil {
		return nil, fmt.Errorf("ProjectRepository.GetFileUrls: %w", err)
	}
	err = db.Select(&dynamicItemUrls, selectQueryForT330, projectID)
	if err != nil {
		return nil, fmt.Errorf("ProjectRepository.GetFileUrls: %w", err)
	}

	fileUrls = append(fileUrls, singleItemUrls...)
	fileUrls = append(fileUrls, multipleItemUrls...)
	fileUrls = append(fileUrls, dynamicItemUrls...)

	return fileUrls, err
}
