package repository_rules

import (
	"fmt"

	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/repository"
)

func (RuleRepository) GetVoicelineLayer(ruleID int) (int, error) {
	db := repository.GetDB()
	var voiceline_layer int
	getQuery := "SELECT voiceline_layer FROM t400_rule WHERE t200_id=?"
	err := db.Get(&voiceline_layer, getQuery, ruleID)
	if err != nil {
		return 0, err
	}
	return voiceline_layer, nil
}

func (RuleRepository) GetSelectedCharacters(ruleID int) ([]models.CharacterItemInRule, error) {
	db := repository.GetDB()
	var characters []models.CharacterItemInRule
	var emptyItems []models.CharacterItemInRule

	selectQueryForT411 := `
	SELECT t110_character.id, t110_character.is_empty,t110_character.name
	FROM t110_character
	INNER JOIN t411_character_item_in_rule
	ON t110_character.id = t411_character_item_in_rule.t110_id
	WHERE t411_character_item_in_rule.t400_id = ?
	`
	err := db.Select(&characters, selectQueryForT411, ruleID)
	if err != nil {
		return nil, err
	}

	selectQueryForT412 := `
	SELECT t110_character.id,  t110_character.name, t110_character.is_empty, t412_empty_item_in_rule.sentence
	FROM t110_character
	INNER JOIN t412_empty_item_in_rule
	ON t110_character.id = t412_empty_item_in_rule.t110_id
	WHERE t412_empty_item_in_rule.t400_id = ?
	`
	err = db.Select(&emptyItems, selectQueryForT412, ruleID)
	if err != nil {
		return nil, err
	}

	characters = append(characters, emptyItems...)
	return characters, nil
}

func (RuleRepository) GetSelectedDynamicItems(ruleID int) ([]models.DynamicItemInRule, error) {
	db := repository.GetDB()
	var dynamicItems []models.DynamicItemInRule
	selectQueryForT420 := `
	SELECT t330_dynamic_item.id, t420_dynamic_item_in_rule.layer
	FROM t330_dynamic_item
	INNER JOIN t420_dynamic_item_in_rule
	ON t330_dynamic_item.id = t420_dynamic_item_in_rule.t330_id
	WHERE t420_dynamic_item_in_rule.t400_id = ?
	`
	err := db.Select(&dynamicItems, selectQueryForT420, ruleID)
	if err != nil {
		return nil, err
	}
	return dynamicItems, nil
}

func (RuleRepository) GetSelectedSingleItems(ruleID int) ([]models.SingleItemInRule, error) {
	db := repository.GetDB() // Get your DB connection

	var selectedSingleItems []models.SingleItemInRule

	selectQuery := `
	SELECT
		t310_single_item.id,
		t430_static_item_in_rule.layer,
		t430_static_item_in_rule.id AS static_item_id,
		t431_single_item_in_rule.is_fixed_start,
		t431_single_item_in_rule.is_fixed_end
	FROM
		t431_single_item_in_rule
		INNER JOIN t310_single_item ON t431_single_item_in_rule.t310_id = t310_single_item.id
		INNER JOIN t430_static_item_in_rule ON t431_single_item_in_rule.t430_id = t430_static_item_in_rule.id
	WHERE
		t430_static_item_in_rule.t400_id = ?
`
	err := db.Select(&selectedSingleItems, selectQuery, ruleID)
	if err != nil {
		return selectedSingleItems, fmt.Errorf("RuleRepository: %v", err)
	}

	// Get Start and End information
	for index, item := range selectedSingleItems {
		if item.IsFixedStart {
			selectQuery = `SELECT insert_place FROM t441_fixed_start WHERE t430_id = ?`
			err = db.Get(&selectedSingleItems[index].Start.InsertPlace, selectQuery, item.StaticItemID)
			if err != nil {
				return selectedSingleItems, fmt.Errorf("RuleRepository.GetSelectedSingleItems: %w", err)
			}
		} else {
			selectQuery = `SELECT t110_id, adjustment_value FROM t442_flexible_start WHERE t430_id = ?`
			err = db.Get(&selectedSingleItems[index].Start, selectQuery, item.StaticItemID)
			if err != nil {
				return selectedSingleItems, fmt.Errorf("RuleRepository.GetSelectedSingleItems: %w", err)
			}
		}

		if item.IsFixedEnd {
			selectQuery = `SELECT length, is_unique FROM t451_fixed_end WHERE t430_id = ?`
			err = db.Get(&selectedSingleItems[index].End, selectQuery, item.StaticItemID)
			if err != nil {
				return selectedSingleItems, fmt.Errorf("RuleRepository.GetSelectedSingleItems: %w", err)
			}
		} else {
			selectQuery = `SELECT how_many_aheads, adjustment_value FROM t452_flexible_end WHERE t430_id = ?`
			err = db.Get(&selectedSingleItems[index].End, selectQuery, item.StaticItemID)
			if err != nil {
				return selectedSingleItems, fmt.Errorf("RuleRepository.GetSelectedSingleItems: %w", err)
			}
		}
	}

	return selectedSingleItems, nil
}

func (RuleRepository) GetSelectedMultipleItems(ruleID int) ([]models.MultipleItemInRule, error) {
	db := repository.GetDB() // Get your DB connection

	var selectedMultipleItems []models.MultipleItemInRule

	selectQuery := `
	SELECT
		t320_multiple_item.id,
		t430_static_item_in_rule.layer,
		t430_static_item_in_rule.id AS static_item_id,
		t432_multiple_item_in_rule.is_fixed_start
	FROM
		t432_multiple_item_in_rule
		INNER JOIN t320_multiple_item ON t432_multiple_item_in_rule.t320_id = t320_multiple_item.id
		INNER JOIN t430_static_item_in_rule ON t432_multiple_item_in_rule.t430_id = t430_static_item_in_rule.id
	WHERE
		t430_static_item_in_rule.t400_id = ?
	`
	err := db.Select(&selectedMultipleItems, selectQuery, ruleID)
	if err != nil {
		return selectedMultipleItems, fmt.Errorf("RuleRepository: %v", err)
	}

	// Get Start information
	for index, item := range selectedMultipleItems {
		if item.IsFixedStart {
			selectQuery = `SELECT insert_place FROM t441_fixed_start WHERE t430_id = ?`
			err = db.Get(&selectedMultipleItems[index].Start.InsertPlace, selectQuery, item.StaticItemID)
			if err != nil {
				return selectedMultipleItems, fmt.Errorf("RuleRepository.GetSelectedMultipleItems: %w", err)
			}
		} else {
			selectQuery = `SELECT t110_id, adjustment_value FROM t442_flexible_start WHERE t430_id = ?`
			err = db.Get(&selectedMultipleItems[index].Start, selectQuery, item.StaticItemID)
			if err != nil {
				return selectedMultipleItems, fmt.Errorf("RuleRepository.GetSelectedMultipleItems: %w", err)
			}
		}
	}

	return selectedMultipleItems, nil
}
