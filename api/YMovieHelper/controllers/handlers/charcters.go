package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/utils"
)

func CreateCharacters(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.CreateCharacters: Invalid file: %v", err)
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.CreateCharacters: Invalid file: %v", err)
		return
	}

	defer openedFile.Close()

	reader, err := utils.UTF8Encoding(openedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.CreateCharacters: %v", err)
		return
	}

	// JSON デコード
	var ymmpData domains.YMMP
	if err := json.NewDecoder(reader).Decode(&ymmpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file."})
		log.Printf("handlers.CreateCharacters: %v", err)
		return
	}

	softwareID, _, _ := utils.GetParams(c)

	model := di_controllers.InitializeCharacterModel()
	err = model.CreateCharacters(&ymmpData, softwareID)
	if err != nil {
		log.Printf("handlers.CreateCharacters: %v", err)
		var customErr *utils.CustomError
		if errors.As(err, &customErr) {
			c.JSON(http.StatusBadRequest, gin.H{"userError": customErr.FrontError()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to create characters."})
}

func GetCharacters(c *gin.Context) {
	softwareID, _, _ := utils.GetParams(c)

	model := di_controllers.InitializeCharacterModel()
	characters, err := model.GetCharacters(softwareID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.GetCharacters: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"characters": characters})
}

func UpdateCharacter(c *gin.Context) {
	var req models.CharacterForUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.UpdateCharacter: %v", err)
		return
	}

	model := di_controllers.InitializeCharacterModel()
	err = model.UpdateCharacter(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UpdateCharacter: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to update software."})
}

func DeleteCharacter(c *gin.Context) {
	_, _, characterID := utils.GetParams(c)

	model := di_controllers.InitializeCharacterModel()
	err := model.DeleteCharacter(c, characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete character"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to delete character."})
}

func CreateCharacterEmotions(c *gin.Context) {
	var ymmpData domains.YMMP
	_, _, characterID := utils.GetParams(c)
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.CreateCharacterEmotions: Invalid file: %v", err)
		return
	}

	defer file.Close()

	reader, err := utils.UTF8Encoding(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.CreateCharacterEmotions: %v", err)
		return
	}

	// JSON デコード
	if err := json.NewDecoder(reader).Decode(&ymmpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file."})
		log.Printf("handlers.CreateCharacterEmotions: %v", err)
		return
	}

	model := di_controllers.InitializeCharacterModel()
	err = model.CreateEmotions(c, &ymmpData, characterID)
	if err != nil {
		log.Printf("handlers.CreateCharacterEmotions: %v", err)
		var customErr *utils.CustomError
		if errors.As(err, &customErr) {
			c.JSON(http.StatusBadRequest, gin.H{"userError": customErr.FrontError()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to create character"})
}

func GetEmotions(c *gin.Context) {
	_, _, characterID := utils.GetParams(c)

	model := di_controllers.InitializeCharacterModel()
	characterName, emotions, err := model.GetEmotions(characterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file."})
		log.Printf("handlers.GetEmotions: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"emotions": emotions, "characterName": characterName})
}

func UpdateEmotion(c *gin.Context) {
	var req models.EmotionForUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.UpdateEmotion: %v", err)
		return
	}

	model := di_controllers.InitializeCharacterModel()
	err = model.UpdateEmotion(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UpdateCharacter: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to update emotion."})
}

func DeleteEmotion(c *gin.Context) {
	type Request struct {
		ID int `json:"id"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.DeleteCharacters: %v", err)
		return
	}

	model := di_controllers.InitializeCharacterModel()
	err := model.DeleteEmotion(c, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete emotion"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to delete emotion."})
}
