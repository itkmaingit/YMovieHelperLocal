package utils

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID string
	jwt.StandardClaims
}

func GetParams(c *gin.Context) (softwareID int, projectID int, characterID int) {

	strSoftwareID := c.Param("softwareID")
	softwareID, err := strconv.Atoi(strSoftwareID)
	if err != nil {
		softwareID = 0
	}

	strProjectID := c.Param("projectID")
	projectID, err = strconv.Atoi(strProjectID)
	if err != nil {
		projectID = 0
	}

	strCharacterID := c.Param("characterID")
	characterID, err = strconv.Atoi(strCharacterID)
	if err != nil {
		characterID = 0
	}

	return softwareID, projectID, characterID
}
