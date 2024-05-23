package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/utils"
)

func DownloadScenarioCSV(c *gin.Context) {
	_, projectID, _ := utils.GetParams(c)

	model := di_controllers.InitializeDownloadModel()
	filePath := model.ConstructDownloadScenarioCSV(projectID)

	c.FileAttachment(filePath, "template.csv")

}
func DownloadScenarioTXT(c *gin.Context) {
	_, projectID, _ := utils.GetParams(c)

	model := di_controllers.InitializeDownloadModel()
	filePath := model.ConstructScenarioTXT(projectID)

	c.FileAttachment(filePath, "scenario.txt")

}
func DownloadCompleteYMMP(c *gin.Context) {
	_, projectID, _ := utils.GetParams(c)

	model := di_controllers.InitializeDownloadModel()
	filePath := model.ConstructDownloadCompleteYMMP(projectID)

	c.FileAttachment(filePath, "complete.ymmp")

}
