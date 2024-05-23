package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/utils"
)

func ResolveSingleItem(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.ResolveSingleItem: Invalid file: %v", err)
		return
	}

	// ファイルを開く
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.ResolveSingleItem: Failed to open file: %v", err)
		return
	}
	defer openedFile.Close()

	reader, err := utils.UTF8Encoding(openedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.ResolveSingleItem: %v", err)
		return
	}

	// JSON デコード
	var ymmpData domains.YMMP
	if err := json.NewDecoder(reader).Decode(&ymmpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file."})
		log.Printf("handlers.ResolveSingleItem: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	descriptions, err := model.GetSingleItemDescriptionsAndSave(c, &ymmpData)
	if err != nil {
		log.Printf("handlers.ResolveSingleItem: %v", err)
		var customErr *utils.CustomError
		if errors.As(err, &customErr) {
			c.JSON(http.StatusBadRequest, gin.H{"userError": customErr.FrontError()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return

	}

	c.JSON(http.StatusOK, gin.H{"singleItems": descriptions})
}

func UploadSingleItem(c *gin.Context) {
	var errors []error
	type Request struct {
		Name     string `json:"name"`
		ID       string `json:"id"`
		ItemType string `json:"itemType"`
		Length   int    `json:"length"`
	}
	type Requests struct {
		Items []Request `json:"items"`
	}
	_, projectID, _ := utils.GetParams(c)

	var req Requests
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.UploadSingleItem: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	for _, req := range req.Items {
		err = model.UploadSingleItem(c, projectID, req.Name, req.ID, req.ItemType, req.Length)
		if err != nil {
			errors = append(errors, fmt.Errorf("handlers.UploadSingleItem: %v", err))
		}
	}

	if len(errors) > 0 {
		errorMessages := make([]string, len(errors))
		for i, err := range errors {
			errorMessages[i] = err.Error()
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UploadSingleItem: multiple errors: \n%s", strings.Join(errorMessages, "; "))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func UploadMultipleItem(c *gin.Context) {
	var ymmpData domains.YMMP
	_, projectID, _ := utils.GetParams(c)
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.UploadMultipleItem: Invalid request.")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.UploadMultipleItem: Invalid file: %v", err)
		return
	}
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.UploadMultipleItem: Invalid file: %v", err)
		return
	}
	defer openedFile.Close()

	reader, err := utils.UTF8Encoding(openedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.UploadMultipleItem: %v", err)
		return
	}

	if err := json.NewDecoder(reader).Decode(&ymmpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file."})
		log.Printf("handlers.UploadMultipleItem: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err = model.UploadMultipleItem(c, &ymmpData, projectID, name)
	if err != nil {
		log.Printf("handlers.UploadMultipleItem: %v", err)
		var customErr *utils.CustomError
		if errors.As(err, &customErr) {
			c.JSON(http.StatusBadRequest, gin.H{"userError": customErr.FrontError()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func UploadDynamicItem(c *gin.Context) {
	var ymmpData domains.YMMP
	_, projectID, _ := utils.GetParams(c)

	name := c.PostForm("name")
	itemType := c.PostForm("itemType")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.UploadDynamicItem: Invalid file: %v", err)
		return
	}
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.UploadDynamicItem: Invalid file: %v", err)
		return
	}
	defer openedFile.Close()

	reader, err := utils.UTF8Encoding(openedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.UploadDynamicItem: %v", err)
		return
	}

	if err := json.NewDecoder(reader).Decode(&ymmpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file."})
		log.Printf("handlers.UploadDynamicItem: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err = model.UploadDynamicItem(c, &ymmpData, projectID, itemType, name)
	if err != nil {
		log.Printf("handlers.UploadDynamicItem: %v", err)
		var customErr *utils.CustomError
		if errors.As(err, &customErr) {
			c.JSON(http.StatusBadRequest, gin.H{"userError": customErr.FrontError()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func GetItems(c *gin.Context) {
	_, projectID, _ := utils.GetParams(c)

	model := di_controllers.InitializeItemModel()
	singleItems, multipleItems, dynamicItems, err := model.GetItems(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.GetItems: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"singleItems": singleItems, "multipleItems": multipleItems, "dynamicItems": dynamicItems})
}

func UpdateSingleItem(c *gin.Context) {
	var req models.SingleItemForUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.UpdateEmotion: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err = model.UpdateSingleItem(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UpdateSingleItem: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to update single item."})
}

func UpdateMultipleItem(c *gin.Context) {
	var req models.MultipleItemForUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.UpdateEmotion: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err = model.UpdateMultipleItem(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UpdateMultipleItem: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to update multiple item."})
}

func UpdateDynamicItem(c *gin.Context) {
	var req models.DynamicItemForUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.UpdateEmotion: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err = model.UpdateDynamicItem(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UpdateDynamicItem: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to update dynamic item."})
}

func DeleteSingleItem(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.DeleteCharacters: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err := model.DeleteSingleItem(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete emotion"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to delete single item."})
}

func DeleteMultipleItem(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.DeleteCharacters: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err := model.DeleteMultipleItem(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete multiple"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to delete single item."})
}

func DeleteDynamicItem(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.DeleteCharacters: %v", err)
		return
	}

	model := di_controllers.InitializeItemModel()
	err := model.DeleteDynamicItem(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete emotion"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to delete dynamic item."})
}
