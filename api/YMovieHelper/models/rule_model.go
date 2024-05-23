package models

import (
	"context"
	"fmt"

	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/jmoiron/sqlx"
)

// TODO: Ruleの中でも、MultipleItem, SingleItem, CharacterItem, DynamicItemで分割
type RuleModel struct {
	ruleRepo  IRuleRepository
	charaRepo ICharacterRepository
	itemRepo  IItemRepository
	csvDomain domains.CSVDomain
}

type IRuleRepository interface {
	ExistsRule(projectID int) (bool, error)
	Create(*RuleForInsert) error
	UpdateRule(projectID int, voicelineLayer int) (*sqlx.Tx, error)

	//InRule化
	CreateCharacterItem(tx *sqlx.Tx, projectID int, characterID int) error
	CreateEmptyItem(tx *sqlx.Tx, projectID int, characterID int, sentence string) error
	CreateDynamicItem(tx *sqlx.Tx, projectID int, layer int, dynamicItemID int) error
	CreateSingleItem(tx *sqlx.Tx, projectID int, singleItem SingleItemInRule) error
	CreateMultipleItem(tx *sqlx.Tx, projectID int, multipleItem MultipleItemInRule) error

	GetVoicelineLayer(projectID int) (int, error)
	GetSelectedCharacters(projectID int) ([]CharacterItemInRule, error)
	GetSelectedDynamicItems(projectID int) ([]DynamicItemInRule, error)
	GetSelectedSingleItems(projectID int) ([]SingleItemInRule, error)
	GetSelectedMultipleItems(projectID int) ([]MultipleItemInRule, error)
	GetCharacterItemIDsInRule(ruleID int) ([]int, error)
	GetEmptyItemIDsInRule(ruleID int) ([]int, error)
}

type RuleForInsert struct {
	T200ID         int `db:"t200_id"`
	VoicelineLayer int `db:"voiceline_layer"`
}

func NewRuleModel(ruleRepo IRuleRepository, charaRepo ICharacterRepository, itemRepo IItemRepository, csvDomain domains.CSVDomain) RuleModel {
	model := RuleModel{ruleRepo: ruleRepo, charaRepo: charaRepo, itemRepo: itemRepo, csvDomain: csvDomain}
	return model
}

func (model RuleModel) CreateRule(ctx context.Context, data *Rule, ruleID int) error {
	//途中、いつでもエラーが起きたら全てロールバックして、returnする
	//まず、ruleのvoiceline_layerをupdate
	// 次に、chara, empty, dynamic, static(multiple, single)のルールの追加
	var tx *sqlx.Tx
	var err error
	tx, err = model.ruleRepo.UpdateRule(ruleID, data.VoicelineLayer)
	if err != nil {
		return fmt.Errorf("RuleModel.CreateRule: %v", err)
	}

	for _, charaItem := range data.CharaItems {
		if charaItem.IsEmpty {
			err = model.ruleRepo.CreateEmptyItem(tx, ruleID, charaItem.CharacterID, charaItem.Sentence)
			if err != nil {
				return fmt.Errorf("RuleModel.CreateRule: %v", err)
			}
		} else if !charaItem.IsEmpty {
			err = model.ruleRepo.CreateCharacterItem(tx, ruleID, charaItem.CharacterID)
			if err != nil {
				return fmt.Errorf("RuleModel.CreateRule: %v", err)
			}
		}
	}

	for _, dynamicItem := range data.DynamicItems {
		err = model.ruleRepo.CreateDynamicItem(tx, ruleID, dynamicItem.Layer, dynamicItem.DynamicItemID)
		if err != nil {
			return fmt.Errorf("RuleModel.CreateRule: %v", err)
		}
	}

	for _, singleItem := range data.SingleItems {
		err = model.ruleRepo.CreateSingleItem(tx, ruleID, singleItem)
		if err != nil {
			return fmt.Errorf("RuleModel.CreateRule: %v", err)
		}
	}

	for _, multipleItem := range data.MultipleItems {
		err = model.ruleRepo.CreateMultipleItem(tx, ruleID, multipleItem)
		if err != nil {
			return fmt.Errorf("RuleModel.CreateRule: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("RuleRepository.UpdateRule: %w", err)
	}

	return err
}

func (model RuleModel) GetRule(softwareID int, projectID int) (RuleResponse, error) {
	var response RuleResponse
	isExists, err := model.ruleRepo.ExistsRule(projectID)
	if err != nil {
		return response, fmt.Errorf("RuleModel.GetRule: %w", err)
	}

	characters, err := model.charaRepo.GetCharacters(softwareID)
	if err != nil {
		return response, fmt.Errorf("RuleModel.GetRule: %w", err)
	}

	dynamicItems, err := model.itemRepo.GetDynamicItems(projectID)
	if err != nil {
		return response, fmt.Errorf("RuleModel.GetRule: %w", err)
	}

	singleItems, err := model.itemRepo.GetSingleItems(projectID)
	if err != nil {
		return response, fmt.Errorf("RuleModel.GetRule: %w", err)
	}

	multipleItems, err := model.itemRepo.GetMultipleItems(projectID)
	if err != nil {
		return response, fmt.Errorf("RuleModel.GetRule: %w", err)
	}

	if isExists {
		voicelineLayer, err := model.ruleRepo.GetVoicelineLayer(projectID)
		if err != nil {
			return response, fmt.Errorf("RuleModel.GetRule: %w", err)
		}
		selectedCharacters, err := model.ruleRepo.GetSelectedCharacters(projectID)
		if err != nil {
			return response, fmt.Errorf("RuleModel.GetRule: %w", err)
		}

		selectedDynamicItems, err := model.ruleRepo.GetSelectedDynamicItems(projectID)
		if err != nil {
			return response, fmt.Errorf("RuleModel.GetRule: %w", err)
		}

		selectedSingleItems, err := model.ruleRepo.GetSelectedSingleItems(projectID)
		if err != nil {
			return response, fmt.Errorf("RuleModel.GetRule: %w", err)
		}

		selectedMultipleItems, err := model.ruleRepo.GetSelectedMultipleItems(projectID)
		if err != nil {
			return response, fmt.Errorf("RuleModel.GetRule: %w", err)
		}

		response = RuleResponse{VoicelineLayer: voicelineLayer, CharaItems: characters, DynamicItems: dynamicItems, SingleItems: singleItems, MultipleItems: multipleItems, SelectedCharacterItems: selectedCharacters, SelectedDynamicItems: selectedDynamicItems, SelectedSingleItems: selectedSingleItems, SelectedMultipleItems: selectedMultipleItems}
		return response, err

	} else {
		response = RuleResponse{CharaItems: characters, DynamicItems: dynamicItems, SingleItems: singleItems, MultipleItems: multipleItems}
		return response, err

	}
}

// アンチパターン、でも面倒すぎるので許容...
type StartInRule struct {
	InsertPlace     string `json:"insertPlace" db:"insert_place"`
	CharacterID     int    `json:"characterID" db:"t110_id"`
	AdjustmentValue int    `json:"adjustmentValue" db:"adjustment_value"`
}

type EndInRule struct {
	IsUnique        bool `json:"isUnique" db:"is_unique"`
	Length          int  `json:"length"`
	HowManyAheads   int  `json:"howManyAheads" db:"how_many_aheads"`
	AdjustmentValue int  `json:"adjustmentValue" db:"adjustment_value"`
}

type CharacterItemInRule struct {
	CharacterID int    `json:"id" db:"id"`
	IsEmpty     bool   `json:"isEmpty" db:"is_empty"`
	Sentence    string `json:"sentence" db:"sentence"`
	Name        string `json:"name" db:"name"`
}

type DynamicItemInRule struct {
	DynamicItemID int `json:"id" db:"id"`
	Layer         int `json:"layer"`
}

type SingleItemInRule struct {
	SingleItemID int         `json:"id" db:"id"`
	StaticItemID int         `db:"static_item_id"`
	Layer        int         `json:"layer" db:"layer"`
	IsFixedStart bool        `json:"isFixedStart" db:"is_fixed_start"`
	Start        StartInRule `json:"start"`
	IsFixedEnd   bool        `json:"isFixedEnd" db:"is_fixed_end"`
	End          EndInRule   `json:"end"`
}

type MultipleItemInRule struct {
	MultipleItemID int         `json:"id" db:"id"`
	StaticItemID   int         `db:"static_item_id"`
	Layer          int         `json:"layer" db:"layer"`
	IsFixedStart   bool        `json:"isFixedStart" db:"is_fixed_start"`
	Start          StartInRule `json:"start"`
}

type Rule struct {
	VoicelineLayer int                   `json:"voicelineLayer"`
	CharaItems     []CharacterItemInRule `json:"charaItems"`
	DynamicItems   []DynamicItemInRule   `json:"dynamicItems"`
	MultipleItems  []MultipleItemInRule  `json:"multipleItems"`
	SingleItems    []SingleItemInRule    `json:"singleItems"`
}

type RuleResponse struct {
	VoicelineLayer         int                     `json:"initialVoicelineLayer"`
	CharaItems             []CharacterForSelect    `json:"charaItems"`
	DynamicItems           []DynamicItemForSelect  `json:"dynamicItems"`
	MultipleItems          []MultipleItemForSelect `json:"multipleItems"`
	SingleItems            []SingleItemForSelect   `json:"singleItems"`
	SelectedCharacterItems []CharacterItemInRule   `json:"initialSelectedCharacterItems"`
	SelectedDynamicItems   []DynamicItemInRule     `json:"initialSelectedDynamicItems"`
	SelectedSingleItems    []SingleItemInRule      `json:"initialSelectedSingleItems"`
	SelectedMultipleItems  []MultipleItemInRule    `json:"initialSelectedMultipleItems"`
}
