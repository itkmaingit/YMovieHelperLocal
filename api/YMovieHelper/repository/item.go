package repository

import (
	"fmt"
	"strings"

	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/utils"
)

type ItemRepository struct {
}

func NewItemRepository() models.IItemRepository {
	return ItemRepository{}
}

func (ItemRepository) GetSingleItems(projectID int) ([]models.SingleItemForSelect, error) {
	db := GetDB()
	var singleItems []models.SingleItemForSelect
	selectQuery := "SELECT id, m301_name, item_path, name,length FROM t310_single_item WHERE t200_id=?"
	err := db.Select(&singleItems, selectQuery, projectID)
	if err != nil {
		return nil, fmt.Errorf("ItemRepository.GetSingleItems: %w", err)
	}
	return singleItems, nil
}

func (ItemRepository) GetMultipleItems(projectID int) ([]models.MultipleItemForSelect, error) {
	db := GetDB()
	var multipleItems []models.MultipleItemForSelect
	selectQuery := "SELECT id, item_path, name,count_of_items FROM t320_multiple_item WHERE t200_id=?"
	err := db.Select(&multipleItems, selectQuery, projectID)
	if err != nil {
		return nil, fmt.Errorf("ItemRepository.GetMultipleItems: %w", err)
	}
	return multipleItems, nil
}

func (ItemRepository) GetDynamicItems(projectID int) ([]models.DynamicItemForSelect, error) {
	db := GetDB()
	var dynamicItems []models.DynamicItemForSelect
	selectQuery := "SELECT id, m301_name, item_url, item_path_in_pc, name FROM t330_dynamic_item WHERE t200_id=?"
	err := db.Select(&dynamicItems, selectQuery, projectID)
	if err != nil {
		return nil, fmt.Errorf("ItemRepository.GetDynamicItems: %w", err)
	}
	return dynamicItems, nil
}

func (ItemRepository) CreateSingleItem(singleItem models.SingleItemForInsert) error {
	db := GetDB()
	var count int
	checkQuery := "SELECT COUNT(*) FROM t310_single_item WHERE t200_id = ?"
	insertSingleItemQuery := "INSERT INTO t310_single_item (t200_id,m301_name,item_path,name,length) VALUES (:t200_id,:m301_name,:item_path,:name,:length)"

	err := db.Get(&count, checkQuery, singleItem.T200ID)
	if err != nil || count >= 10 {
		return fmt.Errorf("ItemRepository.CreateSingleItem: failed to insert data: %w", err)
	}

	_, err = db.NamedExec(insertSingleItemQuery, singleItem)
	if err != nil {
		return fmt.Errorf("ItemRepository.CreateSingleItem: failed to insert data: %w", err)
	}

	return err
}

func (ItemRepository) CreateMultipleItem(multipleItem models.MultipleItemForInsert) error {
	db := GetDB()
	var count int
	checkQuery := "SELECT COUNT(*) FROM t320_multiple_item WHERE t200_id = ?"
	insertMultipleItemQuery := "INSERT INTO t320_multiple_item (t200_id,item_path,name,count_of_items) VALUES (:t200_id,:item_path,:name,:count_of_items)"

	err := db.Get(&count, checkQuery, multipleItem.T200ID)
	if err != nil || count >= 10 {
		return fmt.Errorf("ItemRepository.CreateMultipleItem: failed to insert data: %w", err)
	}

	_, err = db.NamedExec(insertMultipleItemQuery, multipleItem)
	if err != nil {
		return fmt.Errorf("ItemRepository.CreateMultipleItem: failed to insert data: %w", err)
	}

	return err
}

func (ItemRepository) CreateDynamicItem(projectID int, itemType string, fileUrl string, itemPathInPC string, name string) error {
	db := GetDB()
	var count int
	checkQuery := "SELECT COUNT(*) FROM t330_dynamic_item WHERE t200_id = ?"
	insertQuery := "INSERT INTO t330_dynamic_item (t200_id, m301_name, item_url, item_path_in_pc, name) VALUES (?,?,?,?,?)"

	err := db.Get(&count, checkQuery, projectID)
	if err != nil || count >= 10 {
		return fmt.Errorf("ItemRepository.CreateDynamicItem: failed to insert data: %w", err)
	}

	_, err = db.Exec(insertQuery, projectID, itemType, fileUrl, itemPathInPC, name)
	if err != nil {
		return fmt.Errorf("repository.ItemRepository: failed to insert data: %w", err)
	}
	return err
}

func (ItemRepository) GetSingleItemType(typeOnYMMP string) (string, error) {
	db := GetDB()
	var itemType []string
	selectQuery := `SELECT name
									FROM m301_item_type
									WHERE ymmp_name =?`
	err := db.Select(&itemType, selectQuery, typeOnYMMP)
	if err != nil {
		return "", fmt.Errorf("repository.ItemRepository: failed to select item type: %w", err)
	}
	return itemType[0], err

}

func (ItemRepository) UpdateSingleItem(data *models.SingleItemForUpdate) error {
	db := GetDB()
	var updateQuery strings.Builder
	updateQuery.WriteString("UPDATE t310_single_item SET ")

	var args []interface{}
	if data.Name != nil {
		updateQuery.WriteString("name = ?, ")
		args = append(args, *data.Name)
	}

	if data.Length != nil {
		updateQuery.WriteString("length = ?, ")
		args = append(args, *data.Length)
	}

	query := utils.ModifyQuery(updateQuery)
	query += "WHERE id = ?"

	args = append(args, data.ID)

	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("repository.ItemRepository.UpdateSingleItem: %w", err)
	}

	return err
}
func (ItemRepository) UpdateMultipleItem(data *models.MultipleItemForUpdate) error {
	db := GetDB()
	var updateQuery strings.Builder
	updateQuery.WriteString("UPDATE t320_multiple_item SET ")

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
		return fmt.Errorf("repository.ItemRepository.UpdateMultipleItem: %w", err)
	}

	return err
}

func (ItemRepository) UpdateDynamicItem(data *models.DynamicItemForUpdate) error {
	db := GetDB()
	var updateQuery strings.Builder
	updateQuery.WriteString("UPDATE t330_dynamic_item SET ")

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
		return fmt.Errorf("repository.ItemRepository.UpdateDynamicItem: %w", err)
	}

	return err
}

func (ItemRepository) GetFileUrlForSingleItem(itemID int) (string, error) {
	db = GetDB()
	var fileUrl string

	selectQueryForT310 := `
	SELECT t310_single_item.item_path
	FROM t310_single_item
	WHERE t310_single_item.id = ?
`
	err := db.Get(&fileUrl, selectQueryForT310, itemID)
	if err != nil {
		return "", fmt.Errorf("ItemRepository.GetFileUrlsForSingleItem: %w", err)
	}

	return fileUrl, err
}

func (ItemRepository) GetFileUrlForMultipleItem(itemID int) (string, error) {
	db = GetDB()
	var fileUrl string

	selectQueryForT320 := `
	SELECT t320_multiple_item.item_path
	FROM t320_multiple_item
	WHERE t320_multiple_item.id = ?
`
	err := db.Get(&fileUrl, selectQueryForT320, itemID)
	if err != nil {
		return "", fmt.Errorf("ItemRepository.GetFileUrlsForMultipleItem: %w", err)
	}

	return fileUrl, err
}

func (ItemRepository) GetFileUrlForDynamicItem(itemID int) (string, error) {
	db = GetDB()
	var fileUrl string

	selectQueryForT330 := `
	SELECT t330_dynamic_item.item_url
	FROM t330_dynamic_item
	WHERE t330_dynamic_item.id = ?
`
	err := db.Get(&fileUrl, selectQueryForT330, itemID)
	if err != nil {
		return "", fmt.Errorf("ItemRepository.GetFileUrlsForDynamicItem: %w", err)
	}

	return fileUrl, err
}

func (ItemRepository) DeleteSingleItem(itemID int) error {
	db := GetDB()
	deleteQuery := "DELETE FROM t310_single_item WHERE id=?"
	_, err := db.Exec(deleteQuery, itemID)
	if err != nil {
		return fmt.Errorf("ItemRepository.DeleteSingleItem: %w", err)
	}
	return nil
}
func (ItemRepository) DeleteMultipleItem(itemID int) error {
	db := GetDB()
	deleteQuery := "DELETE FROM t320_multiple_item WHERE id=?"
	_, err := db.Exec(deleteQuery, itemID)
	if err != nil {
		return fmt.Errorf("ItemRepository.DeleteMultipleItem: %w", err)
	}
	return nil
}
func (ItemRepository) DeleteDynamicItem(itemID int) error {
	db := GetDB()
	deleteQuery := "DELETE FROM t330_dynamic_item WHERE id=?"
	_, err := db.Exec(deleteQuery, itemID)
	if err != nil {
		return fmt.Errorf("ItemRepository.DeleteDynamicItem: %w", err)
	}
	return nil
}
