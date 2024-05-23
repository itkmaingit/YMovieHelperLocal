package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/models"
	"github.com/itkmaingit/YMovieHelper/utils"
)

type Start struct {
	FixedName       string `json:"fixedName,omitempty"`
	CharacterID     string `json:"characterID,omitempty"`
	AdjustmentValue int    `json:"adjustmentValue,omitempty"`
}

type End struct {
	Length          int    `json:"length,omitempty"`
	HowManyAhead    string `json:"howManyAhead,omitempty"`
	AdjustmentValue int    `json:"adjustmentValue,omitempty"`
}

type VoiceItem struct {
	T110ID int `json:"characterID"`
}

type EmptyItem struct {
	Sentence string `json:"sentence"`
	Name     string `json:"name"`
}

type SingleItem struct {
	T310ID int   `json:"singleItemID"`
	Layer  int   `json:"layer"`
	Start  Start `json:"insert"`
	End    End   `json:"end"`
}

type MultipleItem struct {
	T320ID int `json:"multipleItemID"`
	End    End `json:"end"`
}

type DynamicItem struct {
	M301_Name string `json:"itemType"`
	Layer     int    `json:"layer"`
}

func CreateRule(c *gin.Context) {
	var req models.Rule
	_, projectID, _ := utils.GetParams(c)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.CreateRule: invalid request:%v", err)
		return
	}
	model := di_controllers.InitializeRuleModel()
	err = model.CreateRule(c, &req, projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.CreateRule: invalid request:%v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succeeded rule creation."})
}

func GetRule(c *gin.Context) {
	softwareID, projectID, _ := utils.GetParams(c)
	model := di_controllers.InitializeRuleModel()
	response, err := model.GetRule(softwareID, projectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request."})
		log.Printf("handlers.GetRule: invalid request:%v", err)
		return
	}

	c.JSON(http.StatusOK, response)
}
