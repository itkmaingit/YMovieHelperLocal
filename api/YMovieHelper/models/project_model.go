package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/itkmaingit/YMovieHelper/utils"
)

type ProjectModel struct {
	projectRepo IProjectRepository
}

type IProjectRepository interface {
	Create(*ProjectForInsert) (int, error)
	Update(data *ProjectForUpdate) error
	Delete(projectID int) error

	GetFileUrls(projectID int) ([]string, error)
}

type ProjectForInsert struct {
	T100ID int    `db:"t100_id"`
	Name   string `db:"name"`
}

type ProjectForUpdate struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

func NewProjectModel(projectRepo IProjectRepository) ProjectModel {
	model := ProjectModel{projectRepo: projectRepo}
	return model
}

func (model ProjectModel) CreateProject(data *ProjectForInsert) error {
	_, err := model.projectRepo.Create(data)
	if err != nil {
		return fmt.Errorf("ProjectModel.CreateProject: failed to create project: %w", err)
	}
	return err
}

func (model ProjectModel) UpdateProject(data *ProjectForUpdate) error {
	err := model.projectRepo.Update(data)
	if err != nil {
		return fmt.Errorf("Model.UpdateProject: %w", err)
	}

	return nil
}

func (model ProjectModel) DeleteProject(ctx context.Context, projectID int) error {
	var errorMessages []error
	fileUrls, err := model.projectRepo.GetFileUrls(projectID)
	if err != nil {
		return fmt.Errorf("ProjectModel.DeleteProject: %w", err)
	}
	err = model.projectRepo.Delete(projectID)
	if err != nil {
		return fmt.Errorf("ProjectModel.DeleteProject: %w", err)
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
			errorMessage.WriteString(fmt.Sprintf("ProjectModel.DeleteProject: %v\n", err))
		}
		return errors.New(errorMessage.String())
	}

	return err
}
