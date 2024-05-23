package repository_rules

import (
	"fmt"

	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/repository"
	"github.com/jmoiron/sqlx"
)

type RuleRepository struct {
}

func NewRuleRepository() models.IRuleRepository {
	return RuleRepository{}
}

func NewMakeYMMPRepository() models.IMakeYMMPRepository {
	return RuleRepository{}
}

func NewRuleRepositoryForDomain() domains.IRuleRepository {
	return RuleRepository{}
}

func (RuleRepository) ExistsRule(projectID int) (bool, error) {
	db := repository.GetDB()
	var count int
	selectQuery := "SELECT COUNT(*) FROM t400_rule WHERE t200_id =?"
	err := db.Get(&count, selectQuery, projectID)
	if err != nil {
		return false, fmt.Errorf("RuleRepository,ExistsRule: %w", err)
	}
	return count > 0, nil
}

func (RuleRepository) ExistsCharacter(ruleID int) (bool, error) {
	db := repository.GetDB()
	var count int
	selectQuery := "SELECT COUNT(*) FROM t411_character_item_in_rule WHERE t400_id =?"
	err := db.Get(&count, selectQuery, ruleID)
	if err != nil {
		return false, fmt.Errorf("RuleRepository,ExistsCharacter: %w", err)
	}
	return count > 0, nil
}

func (RuleRepository) Create(data *models.RuleForInsert) error {
	db := repository.GetDB()
	insertQuery := "INSERT INTO t400_rule (t200_id, voiceline_layer) VALUES (:t200_id,:voiceline_layer)"
	_, err := db.NamedExec(insertQuery, *data)
	if err != nil {
		return fmt.Errorf("RuleRepository.Create: %v", err)
	}
	return err
}

// 技術的負債...
// ルールを新しくリクエストでもらうたびに、ルールが存在すればそのルールを削除して新しく作り直す
// ただし、失敗した場合はロールバックを行う
func (RuleRepository) UpdateRule(projectID int, voicelineLayer int) (*sqlx.Tx, error) {
	db := repository.GetDB()
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	deleteQuery := "DELETE FROM t400_rule WHERE t200_id = ?"
	_, err = tx.Exec(deleteQuery, projectID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("RuleRepository.UpdateRule: %w", err)
	}
	insertQuery := "INSERT INTO t400_rule (t200_id, voiceline_layer) VALUES(?, ?)"

	_, err = tx.Exec(insertQuery, projectID, voicelineLayer)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("RuleRepository.UpdateRule: %w", err)
	}

	return tx, err
}

func (RuleRepository) CreateCharacterItem(tx *sqlx.Tx, projectID int, characterID int) error {
	if tx != nil {
		insertQuery := "INSERT INTO t411_character_item_in_rule (t400_id,t110_id) VALUES (?,?)"
		_, err := tx.Exec(insertQuery, projectID, characterID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateCharacterItem: %w", err)
		}
		return err
	}

	return nil
}

func (RuleRepository) CreateEmptyItem(tx *sqlx.Tx, projectID int, characterID int, sentence string) error {
	if tx != nil {
		insertQuery := "INSERT INTO t412_empty_item_in_rule (t400_id,t110_id,sentence) VALUES (?,?,?)"
		_, err := tx.Exec(insertQuery, projectID, characterID, sentence)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateEmptyItem: %w", err)
		}
		return err
	}

	return nil
}

func (RuleRepository) CreateDynamicItem(tx *sqlx.Tx, projectID int, layer int, dynamicItemID int) error {
	if tx != nil {
		insertQuery := "INSERT INTO t420_dynamic_item_in_rule (t400_id,layer,t330_id) VALUES (?,?,?)"
		_, err := tx.Exec(insertQuery, projectID, layer, dynamicItemID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateDynamicItem: %w", err)
		}
		return err
	}

	return nil
}

func (RuleRepository) CreateSingleItem(tx *sqlx.Tx, projectID int, singleItem models.SingleItemInRule) error {
	if tx != nil {
		insertQueryForT430 := "INSERT INTO t430_static_item_in_rule (t400_id,layer) VALUES (?,?)"
		insertQueryForT431 := "INSERT INTO t431_single_item_in_rule (t430_id, t310_id,is_fixed_start,is_fixed_end) VALUES (?,?,?,?)"
		insertQueryForT441 := "INSERT INTO t441_fixed_start (t430_id, insert_place) VALUES (?,?)"
		insertQueryForT442 := "INSERT INTO t442_flexible_start (t430_id, t110_id,adjustment_value) VALUES (?,?,?)"
		insertQueryForT451 := "INSERT INTO t451_fixed_end (t430_id,length,is_unique) VALUES (?,?,?)"
		insertQueryForT452 := "INSERT INTO t452_flexible_end (t430_id, how_many_aheads, adjustment_value) VALUES (?,?,?)"

		result, err := tx.Exec(insertQueryForT430, projectID, singleItem.Layer)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
		}

		staticItemID, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
		}

		_, err = tx.Exec(insertQueryForT431, staticItemID, singleItem.SingleItemID, singleItem.IsFixedStart, singleItem.IsFixedEnd)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
		}

		//T441,
		if singleItem.IsFixedStart {
			_, err = tx.Exec(insertQueryForT441, staticItemID, singleItem.Start.InsertPlace)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
			}
			//  T451
			if singleItem.IsFixedEnd {
				_, err = tx.Exec(insertQueryForT451, staticItemID, singleItem.End.Length, singleItem.End.IsUnique)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
				}

				//T452
			} else if !singleItem.IsFixedEnd {
				_, err = tx.Exec(insertQueryForT452, staticItemID, singleItem.End.HowManyAheads, singleItem.End.AdjustmentValue)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
				}
			}
			//T442
		} else if !singleItem.IsFixedStart {
			_, err = tx.Exec(insertQueryForT442, staticItemID, singleItem.Start.CharacterID, singleItem.Start.AdjustmentValue)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
			}
			//T451
			if singleItem.IsFixedEnd {
				_, err = tx.Exec(insertQueryForT451, staticItemID, singleItem.End.Length, singleItem.End.IsUnique)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
				}
				//T452
			} else if !singleItem.IsFixedEnd {
				_, err = tx.Exec(insertQueryForT452, staticItemID, singleItem.End.HowManyAheads, singleItem.End.AdjustmentValue)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
				}
			}

		}
		return err
	}
	return nil
}

func (RuleRepository) CreateMultipleItem(tx *sqlx.Tx, projectID int, multipleItem models.MultipleItemInRule) error {
	if tx != nil {
		insertQueryForT430 := "INSERT INTO t430_static_item_in_rule (t400_id,layer) VALUES (?,?)"
		insertQueryForT432 := "INSERT INTO t432_multiple_item_in_rule (t430_id, t320_id,is_fixed_start) VALUES (?,?,?)"
		insertQueryForT441 := "INSERT INTO t441_fixed_start (t430_id, insert_place) VALUES (?,?)"
		insertQueryForT442 := "INSERT INTO t442_flexible_start (t430_id, t110_id,adjustment_value) VALUES (?,?,?)"

		result, err := tx.Exec(insertQueryForT430, projectID, multipleItem.Layer)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
		}

		staticItemID, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
		}

		_, err = tx.Exec(insertQueryForT432, staticItemID, multipleItem.MultipleItemID, multipleItem.IsFixedStart)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
		}
		//T441
		if multipleItem.IsFixedStart {
			_, err = tx.Exec(insertQueryForT441, staticItemID, multipleItem.Start.InsertPlace)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
			}

			//T442
		} else if !multipleItem.IsFixedStart {
			_, err = tx.Exec(insertQueryForT442, staticItemID, multipleItem.Start.CharacterID, multipleItem.Start.AdjustmentValue)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("RuleRepository.CreateSingleItem: %w", err)
			}
		}

		return err
	}
	return nil
}

func (RuleRepository) GetDynamicItemNames(ruleID int) ([]string, error) {
	var names []string
	db := repository.GetDB()
	selectQuery := `
	SELECT t330_dynamic_item.name
	FROM t420_dynamic_item_in_rule
	INNER JOIN t330_dynamic_item
	ON t420_dynamic_item_in_rule.t330_id=t330_dynamic_item.id
	WHERE t400_id=?`
	err := db.Select(&names, selectQuery, ruleID)
	if err != nil {
		return names, fmt.Errorf("RuleRepository.GetDynamicItemNames: %w", err)
	}

	return names, nil
}

func (RuleRepository) GetCharacterItemIDsInRule(ruleID int) ([]int, error) {
	var characterIDs []int
	db := repository.GetDB()
	selectQuery := `
	SELECT t110_id FROM t411_character_item_in_rule WHERE t400_id =?
	`

	err := db.Select(&characterIDs, selectQuery, ruleID)
	if err != nil {
		return characterIDs, fmt.Errorf("RuleRepository.GetCharacterItemIDsInRule: %w", err)
	}

	return characterIDs, nil
}

func (RuleRepository) GetEmptyItemIDsInRule(ruleID int) ([]int, error) {
	var emptyItemIDs []int
	db := repository.GetDB()
	selectQuery := `
	SELECT t110_id FROM t412_empty_item_in_rule WHERE t400_id =?
	`

	err := db.Select(&emptyItemIDs, selectQuery, ruleID)
	if err != nil {
		return emptyItemIDs, fmt.Errorf("RuleRepository.GetCharacterItemIDsInRule: %w", err)
	}

	return emptyItemIDs, nil
}
