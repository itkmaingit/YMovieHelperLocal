package middlewares

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/di_containers/di_controllers"
	"github.com/itkmaingit/YMovieHelper/repository"
	"github.com/itkmaingit/YMovieHelper/utils"
)

func AccessControllMiddleware() gin.HandlerFunc {
	// JWTトークンが正当なら、uidなど、必要な情報を全てcontextに結び付けておく
	// パスパラメータがあるリクエストを受けた時、それをuidから照合して、アクセスできないコンテンツならmypageにリダイレクトする
	return func(c *gin.Context) {
		jwtKey := []byte(os.Getenv("JWTKey"))
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims := &utils.Claims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			log.Printf("middlewares.AccessControllMiddleware: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			return
		}

		c.Set("uid", claims.UserID)

		strSoftwareID := c.Param("softwareID")
		strProjectID := c.Param("projectID")
		strCharacterID := c.Param("characterID")

		if strSoftwareID != "" {
			softwareID, err := strconv.Atoi(strSoftwareID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "invalid softwareID",
				})
				log.Printf("invalid softwareID : %v", err)
				return
			}
			softwareModel := di_controllers.InitializeSoftwareModel()
			exists, err := softwareModel.ExistsSoftware(softwareID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed to check if software exists.",
				})
				log.Printf("middlewares.AccessControllMiddleware : %v", err)
				return
			}
			if !exists {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Invalid access.",
				})
				log.Printf("Invalid access.\n")
				return
			}
		}

		if strProjectID != "" {
			if strSoftwareID == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Invalid access",
				})
				log.Printf("middlewares.AccessControllMiddleware:Invalid access.\n")
				return
			}

			softwareID, _ := strconv.Atoi(strSoftwareID)
			projectID, err := strconv.Atoi(strProjectID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "invalid projectID",
				})
				log.Printf("invalid projectID : %v", err)
				return
			}

			repo := repository.ProjectRepository{}
			exists, err := repo.ExistsProject(claims.UserID, softwareID, projectID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed to check if project exists.",
				})
				log.Printf("middlewares.AccessControllMiddleware: Failed to check if project exists : %v", err)
				return
			}

			if !exists {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Invalid access",
				})
				log.Printf("Invalid access.\n")
				return
			}
		}

		if strCharacterID != "" {
			if strSoftwareID == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Invalid access",
				})
				log.Printf("middlewares.AccessControllMiddleware:Invalid access.\n")
				return
			}

			softwareID, _ := strconv.Atoi(strSoftwareID)
			characterID, err := strconv.Atoi(strCharacterID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "invalid projectID",
				})
				log.Printf("invalid projectID : %v", err)
				return
			}

			repo := repository.CharacterRepository{}
			exists, err := repo.ExistsCharacter(softwareID, characterID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "Failed to check if project exists.",
				})
				log.Printf("middlewares.AccessControllMiddleware: Failed to check if project exists : %v", err)
				return
			}

			if !exists {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"message": "Invalid access",
				})
				log.Printf("Invalid access.\n")
				return
			}
		}

		c.Next()
	}
}
