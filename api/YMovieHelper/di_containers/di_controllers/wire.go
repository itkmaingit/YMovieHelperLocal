//go:build wireinject
// +build wireinject

package di_controllers

import (
	"github.com/google/wire"
	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/repository"
	"github.com/itkmaingit/YMovieHelper/repository/repository_rules"
)

func InitializeSoftwareModel() models.SoftwareModel {
	wire.Build(models.NewSoftwareModel, repository.NewSoftwareRepository, repository.NewProjectRepository)
	return models.SoftwareModel{}
}

func InitializeProjectModel() models.ProjectModel {
	wire.Build(models.NewProjectModel, repository.NewProjectRepository)
	return models.ProjectModel{}
}

func InitializeItemModel() models.ItemModel {
	wire.Build(models.NewItemModel, repository.NewItemRepository)
	return models.ItemModel{}
}

func InitializeCharacterModel() models.CharacterModel {
	wire.Build(models.NewCharacterModel, repository.NewCharacterRepository)
	return models.CharacterModel{}
}

func InitializeRuleModel() models.RuleModel {
	wire.Build(models.NewRuleModel,
		repository_rules.NewRuleRepository,
		repository_rules.NewRuleRepositoryForDomain,
		repository.NewCharacterRepository,
		repository.NewItemRepository,
		domains.NewCSVDomain)
	return models.RuleModel{}
}

func InitializeMakeYMMPModel() models.MakeYMMPModel {
	wire.Build(models.NewMakeYMMPModel,
		repository.NewCharacterRepository,
		repository_rules.NewMakeYMMPRepository,
		repository_rules.NewRuleRepository,
		domains.NewCSVDomain,
		repository_rules.NewRuleRepositoryForDomain)
	return models.MakeYMMPModel{}
}
func InitializeDownloadModel() models.DownloadModel {
	wire.Build(models.NewDownloadModel)
	return models.DownloadModel{}
}
