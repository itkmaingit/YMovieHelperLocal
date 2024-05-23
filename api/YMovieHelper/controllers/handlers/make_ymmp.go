package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/models/domains"
	"github.com/itkmaingit/YMovieHelper/utils"
)

var UserRates = make(map[string]*UserRate)

type UserRate struct {
	Count int
	Timer time.Time
}

func CheckRules(c *gin.Context) {
	softwareID, projectID, _ := utils.GetParams(c)

	model := di_controllers.InitializeMakeYMMPModel()
	canMakeYMMP, err := model.CheckRules(c, softwareID, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		log.Printf("handlers.CheckRules: %v", err)
		return
	}

	if !canMakeYMMP {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})

}

func ResolveScenario(c *gin.Context) {
	_, projectID, _ := utils.GetParams(c)
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request."})
		log.Printf("handlers.Scenario, failed to get file.: %v", err)
		return
	}
	defer file.Close()

	model := di_controllers.InitializeMakeYMMPModel()
	fileUrl, err := model.ResolveScenario(c, file, projectID)
	if err != nil {
		log.Printf("handlers.ResolveScenario: %v", err)
		var customErr *utils.CustomError
		if errors.As(err, &customErr) {
			c.JSON(http.StatusBadRequest, gin.H{"userError": customErr.FrontError()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fileUrl": fileUrl})
}

func MakeYMMP(c *gin.Context) {
	_, projectID, _ := utils.GetParams(c)

	movieName := c.PostForm("movieName")

	var ymmpData domains.YMMP
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request."})
		log.Printf("handlers.MakeYMMP: %v", err)
		return
	}

	defer file.Close()

	reader, err := utils.UTF8Encoding(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file"})
		log.Printf("handlers.MakeYMMP: %v", err)
		return
	}

	if err := json.NewDecoder(reader).Decode(&ymmpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid file."})
		log.Printf("handlers.MakeYMMP: %v", err)
		return
	}

	model := di_controllers.InitializeMakeYMMPModel()

	fileUrl, err := model.MakeYMMP(c, &ymmpData, projectID, movieName)
	if err != nil {
		log.Printf("handlers.MakeYMMP: %v", err)
		var customErr *utils.CustomError
		if errors.As(err, &customErr) {
			c.JSON(http.StatusBadRequest, gin.H{"userError": customErr.FrontError()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fileUrl": fileUrl})
}
