package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/utils"
)

func GetSoftwaresAndProjects(c *gin.Context) {

	model := di_controllers.InitializeSoftwareModel()
	softwares, err := model.GetSoftwaresAndProjects()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.GetSoftwaresAndProjects : %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"softwares": softwares})
}

//	@Summary		ソフトウェアデータのアップロード
//	@Description	ソフトウェアデータのアップロード
//	@Tags			users
//	@Accept			json
//	@Produce		plain
//	@Param			user	body		models.User	true	"User"
//	@Success		200		{string}	string		"Succeeded User Registration!"
//	@Router			/upload/software [post]

func CreateSoftware(c *gin.Context) {
	type Request struct {
		Name string `json:"softwareName"`
	}
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.CreateSoftware: %v", err)
		return
	}
	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.CreateSoftware: Invalid parameters.\n")
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "There is not your id. Please signup again.",
		})
		log.Printf("handlers.CreateSoftware: %v\n", err)
	}

	softwareModel := di_controllers.InitializeSoftwareModel()
	err = softwareModel.CreateSoftware(&models.SoftwareForInsert{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.CreateSoftware: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to create software."})
}

func UpdateSoftware(c *gin.Context) {
	var req models.SoftwareForUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Parameters"})
		log.Printf("handlers.UpdateSoftware: %v", err)
		return
	}

	model := di_controllers.InitializeSoftwareModel()
	err = model.UpdateSoftware(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.UpdateSoftware: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to update software."})
}

func DeleteSoftware(c *gin.Context) {
	softwareID, _, _ := utils.GetParams(c)

	model := di_controllers.InitializeSoftwareModel()
	err := model.DeleteSoftware(c, softwareID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error."})
		log.Printf("handlers.DeleteSoftware: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Succeeded to delete software."})
}
