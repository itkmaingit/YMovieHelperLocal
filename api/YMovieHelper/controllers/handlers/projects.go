package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/utils"
)

func CreateProject(c *gin.Context) {
	type Request struct {
		Name string `json:"name"`
	}
	var req Request
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.CreateProject: %v", err)
		return
	}
	softwareID, _, _ := utils.GetParams(c)

	data := models.ProjectForInsert{T100ID: softwareID, Name: req.Name}
	projectModel := di_controllers.InitializeProjectModel()
	err = projectModel.CreateProject(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.CreateProject: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to create project"})
}

func UpdateProject(c *gin.Context) {
	var req models.ProjectForUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.UpdateProject: %v", err)
		return
	}

	model := di_controllers.InitializeProjectModel()
	err = model.UpdateProject(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UpdateProject: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to update project."})
}

func DeleteProject(c *gin.Context) {
	_, projectID, _ := utils.GetParams(c)

	model := di_controllers.InitializeProjectModel()
	err := model.DeleteProject(c, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})

		log.Printf("handlers.DeleteProject: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Suceeded to delete project."})
}
