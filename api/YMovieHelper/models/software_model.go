package models

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/itkmaingit/YMovieHelper/utils"
)

type ISoftwareRepository interface {
	Create(*SoftwareForInsert) (int, error)
	Delete(softwareID int) error
	Update(data *SoftwareForUpdate) error
	GetSoftwares() ([]SoftwareAndProjects, error)
	GetProjects(softwareID int) ([]ProjectForSelect, error)
	ExistsSoftware(softwareID int) (bool, error)

	GetFileUrls(softwareID int) ([]string, error)
}

type SoftwareAndProjects struct {
	SoftwareID int                `db:"id" json:"softwareID"`
	Name       string             `db:"name" json:"name"`
	Projects   []ProjectForSelect `json:"projects"`
}

type ProjectForSelect struct {
	ProjectID int    `db:"id" json:"projectID"`
	Name      string `db:"name" json:"name"`
}

type SoftwareModel struct {
	softwareRepo ISoftwareRepository
	projectRepo  IProjectRepository
}

type SoftwareForInsert struct {
	Name string `db:"name"`
}

type SoftwareForUpdate struct {
	ID   int     `json:"id"`
	Name *string `json:"name"`
}

func NewSoftwareModel(softwareRepo ISoftwareRepository, projectRepo IProjectRepository) SoftwareModel {
	model := SoftwareModel{softwareRepo: softwareRepo, projectRepo: projectRepo}
	return model
}

func (model SoftwareModel) CreateSoftware(softwareData *SoftwareForInsert) error {
	softwareID, err := model.softwareRepo.Create(softwareData)
	if err != nil {
		return fmt.Errorf("Model.CreateSoftware: %w", err)
	}

	projectData := &ProjectForInsert{T100ID: softwareID, Name: "Your Project"}
	_, err = model.projectRepo.Create(projectData)
	if err != nil {
		return fmt.Errorf("Model.CreateSoftware: %w", err)
	}

	return err
}

func (model SoftwareModel) UpdateSoftware(data *SoftwareForUpdate) error {
	err := model.softwareRepo.Update(data)
	if err != nil {
		return fmt.Errorf("Model.UpdateSoftware: %w", err)
	}

	return nil
}

func (model SoftwareModel) DeleteSoftware(ctx context.Context, softwareID int) error {
	var errorMessages []error
	fileUrls, err := model.softwareRepo.GetFileUrls(softwareID)
	if err != nil {
		return fmt.Errorf("SoftwareModel.DeleteSoftware: %w", err)
	}

	err = model.softwareRepo.Delete(softwareID)
	if err != nil {
		return fmt.Errorf("SoftwareModel.DeleteSoftware: %w", err)
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
			errorMessage.WriteString(fmt.Sprintf("SoftwareModel.DeleteSoftware: %v\n", err))
		}
		return errors.New(errorMessage.String())
	}

	return nil
}

func (model SoftwareModel) GetSoftwaresAndProjects() ([]SoftwareAndProjects, error) {
	var projects []ProjectForSelect
	softwares, err := model.softwareRepo.GetSoftwares()
	if err != nil {
		return nil, fmt.Errorf("Model.ReadAllSoftwares: %w", err)
	}
	for i, value := range softwares {
		projects, err = model.softwareRepo.GetProjects(value.SoftwareID)

		if err != nil {
			return nil, fmt.Errorf("Model.ReadAllProjects: %w", err)
		} else {
			softwares[i].Projects = projects
		}
	}

	return softwares, nil
}

func (model SoftwareModel) ExistsSoftware(softwareID int) (bool, error) {
	exists, err := model.softwareRepo.ExistsSoftware(softwareID)
	if err != nil {
		return exists, fmt.Errorf("models.SoftwareModel.ExistsSoftware: %w", err)
	}
	return exists, err
}
