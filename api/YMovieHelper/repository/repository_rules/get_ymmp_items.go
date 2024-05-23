package repository_rules

import (
	"fmt"

	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/repository"
	"github.com/itkmaingit/YMovieHelper/utils"
)

func (RuleRepository) GetDynamicItemsInRule(ruleID int) ([]domains.DynamicItem, error) {
	db := repository.GetDB()
	selectQuery := `
	SELECT t330_dynamic_item.item_url, t330_dynamic_item.item_path_in_pc, t330_dynamic_item.name, t420_dynamic_item_in_rule.layer
	FROM t330_dynamic_item
	INNER JOIN t420_dynamic_item_in_rule
	ON t330_dynamic_item.id = t420_dynamic_item_in_rule.t330_id
	WHERE t420_dynamic_item_in_rule.t400_id=?`
	var dynamicItems []domains.DynamicItem

	err := db.Select(&dynamicItems, selectQuery, ruleID)
	if err != nil {
		return nil, fmt.Errorf("RuleRepository.GetDynamicItemsInRule: %v", err)
	}

	for index, dynamicItem := range dynamicItems {
		dynamicItems[index].ItemPath, err = utils.Decrypt(dynamicItem.ItemPath)
		if err != nil {
			return nil, fmt.Errorf("RuleRepository.GetDynamicItemsInRule: %v", err)
		}
	}

	return dynamicItems, nil
}

func (RuleRepository) GetSingleItemsInRule(ruleId int) ([]domains.SingleItem, error) {
	db := repository.GetDB()
	selectQuery := `
	SELECT t430_static_item_in_rule.id, t310_single_item.item_path, t430_static_item_in_rule.layer, t431_single_item_in_rule.is_fixed_start, t431_single_item_in_rule.is_fixed_end
	FROM t310_single_item
	INNER JOIN t431_single_item_in_rule
	ON t310_single_item.id = t431_single_item_in_rule.t310_id
	INNER JOIN t430_static_item_in_rule
	ON t430_static_item_in_rule.id = t431_single_item_in_rule.t430_id
	WHERE t430_static_item_in_rule.t400_id=?
	`

	selectQueryForFixedStart := `
	SELECT  t441_fixed_start.insert_place
	FROM t441_fixed_start
	WHERE t441_fixed_start.t430_id=?
	`

	selectQueryForFlexibleStart := `
	SELECT  t442_flexible_start.adjustment_value, t110_character.name
	FROM t442_flexible_start
	INNER JOIN t110_character
	ON t442_flexible_start.t110_id = t110_character.id
	WHERE t442_flexible_start.t430_id=?
	`

	selectQueryForFixedEnd := `
	SELECT  t451_fixed_end.length
	FROM t451_fixed_end
	WHERE t451_fixed_end.t430_id=?
	`

	selectQueryForFlexibleEnd := `
	SELECT  t452_flexible_end.how_many_aheads, t452_flexible_end.adjustment_value
	FROM t452_flexible_end
	WHERE t452_flexible_end.t430_id=?
	`

	var singleItems []domains.SingleItem

	err := db.Select(&singleItems, selectQuery, ruleId)
	if err != nil {
		return nil, fmt.Errorf("RuleRepository.GetSingleItemsInRule: %w", err)
	}

	for index, item := range singleItems {
		if item.IsFixedStart {
			err := db.Get(&singleItems[index].Start, selectQueryForFixedStart, item.StaticItemID)
			if err != nil {
				return nil, fmt.Errorf("RuleRepository.GetSingleItemsInRule: %w", err)
			}
		} else if !item.IsFixedStart {
			err := db.Get(&singleItems[index].Start, selectQueryForFlexibleStart, item.StaticItemID)
			if err != nil {
				return nil, fmt.Errorf("RuleRepository.GetSingleItemsInRule: %w", err)
			}
		}

		if item.IsFixedEnd {
			err := db.Get(&singleItems[index].End, selectQueryForFixedEnd, item.StaticItemID)
			if err != nil {
				return nil, fmt.Errorf("RuleRepository.GetSingleItemsInRule: %w", err)
			}
		} else if !item.IsFixedEnd {
			err := db.Get(&singleItems[index].End, selectQueryForFlexibleEnd, item.StaticItemID)
			if err != nil {
				return nil, fmt.Errorf("RuleRepository.GetSingleItemsInRule: %w", err)
			}
		}
	}

	return singleItems, err
}

func (RuleRepository) GetMultipleItemsInRule(ruleId int) ([]domains.MultipleItem, error) {
	db := repository.GetDB()
	selectQuery := `
	SELECT t430_static_item_in_rule.id, t320_multiple_item.item_path, t430_static_item_in_rule.layer, t432_multiple_item_in_rule.is_fixed_start
	FROM t320_multiple_item
	INNER JOIN t432_multiple_item_in_rule
	ON t320_multiple_item.id = t432_multiple_item_in_rule.t320_id
	INNER JOIN t430_static_item_in_rule
	ON t430_static_item_in_rule.id = t432_multiple_item_in_rule.t430_id
	WHERE t430_static_item_in_rule.t400_id=?
	`

	selectQueryForFixedStart := `
	SELECT  t441_fixed_start.insert_place
	FROM t441_fixed_start
	WHERE t441_fixed_start.t430_id=?
	`

	selectQueryForFlexibleStart := `
	SELECT  t442_flexible_start.adjustment_value, t110_character.name
	FROM t442_flexible_start
	INNER JOIN t110_character
	ON t442_flexible_start.t110_id = t110_character.id
	WHERE t442_flexible_start.t430_id=?
	`

	var multipleItems []domains.MultipleItem

	err := db.Select(&multipleItems, selectQuery, ruleId)
	if err != nil {
		return nil, fmt.Errorf("RuleRepository.GetMultipleItemsInRule: %w", err)
	}

	for index, item := range multipleItems {
		if item.IsFixedStart {
			err := db.Get(&multipleItems[index].Start, selectQueryForFixedStart, item.StaticItemID)
			if err != nil {
				return nil, fmt.Errorf("RuleRepository.GetMultipleItemsInRule: %w", err)
			}
		} else if !item.IsFixedStart {
			err := db.Get(&multipleItems[index].Start, selectQueryForFlexibleStart, item.StaticItemID)
			if err != nil {
				return nil, fmt.Errorf("RuleRepository.GetMultipleItemsInRule: %w", err)
			}
		}
	}

	return multipleItems, err
}

func (RuleRepository) GetEmptyItems(ruleID int) ([]domains.EmptyItem, error) {
	db := repository.GetDB()
	var emptyItems []domains.EmptyItem
	selectQuery := `
	SELECT t110_character.name, t412_empty_item_in_rule.sentence
	FROM t110_character
	INNER JOIN t412_empty_item_in_rule
	ON t110_character.id = t412_empty_item_in_rule.t110_id
	WHERE t412_empty_item_in_rule.t400_id=?
	`

	err := db.Select(&emptyItems, selectQuery, ruleID)
	if err != nil {
		return nil, fmt.Errorf("RuleRepository.GetEmptyItems: %w", err)
	}

	return emptyItems, nil
}
func (RuleRepository) GetEmotionMap(ruleID int) (map[string]map[string]string, error) {
	db := repository.GetDB()
	emotionMap := make(map[string]map[string]string)
	selectQuery := `
	SELECT t110_character.name, t110_character.id
	FROM t110_character
	INNER JOIN t411_character_item_in_rule
	ON t110_character.id = t411_character_item_in_rule.t110_id
	WHERE t411_character_item_in_rule.t400_id=?
	`

	selectQueryForEmotion := `
	SELECT t111_emotion.name,t111_emotion.item_path
	FROM t111_emotion
	WHERE t111_emotion.t110_id=?
	`

	rows, err := db.Queryx(selectQuery, ruleID)
	if err != nil {
		return nil, fmt.Errorf("RuleRepository.GetEmotionMap: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var characterName string
		var characterID int
		err := rows.Scan(&characterName, &characterID)
		if err != nil {
			return nil, fmt.Errorf("RuleRepository.GetEmotionMap: %w", err)
		}

		innerMap, ok := emotionMap[characterName]
		if !ok {
			innerMap = make(map[string]string)
			emotionMap[characterName] = innerMap
		}
		emotionRows, err := db.Queryx(selectQueryForEmotion, characterID)
		if err != nil {
			return nil, fmt.Errorf("RuleRepository.GetEmotionMap: %w", err)
		}
		for emotionRows.Next() {
			var emotionName string
			var emotionItemPath string
			err := emotionRows.Scan(&emotionName, &emotionItemPath)
			if err != nil {
				return nil, fmt.Errorf("RuleRepository.GetEmotionMap: %w", err)
			}
			innerMap[emotionName] = emotionItemPath
		}
		emotionRows.Close()
	}

	return emotionMap, err

}

func (RuleRepository) GetCharacterNamesInRule(ruleID int) ([]string, error) {
	db := repository.GetDB()
	var characterNames []string
	selectQuery := `
	SELECT t110_character.name
	FROM t110_character
	INNER JOIN t411_character_item_in_rule
	ON t110_character.id = t411_character_item_in_rule.t110_id
	WHERE t411_character_item_in_rule.t400_id=?
	`

	err := db.Select(&characterNames, selectQuery, ruleID)
	if err != nil {
		return nil, fmt.Errorf("RuleRepository.GetCharacterNamesInRule: %w", err)
	}

	return characterNames, err
}
