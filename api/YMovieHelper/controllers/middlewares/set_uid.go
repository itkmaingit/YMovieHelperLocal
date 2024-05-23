package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request."})
			c.Abort()
			return
		}

		c.Set("uid", uid)
		c.Next()
	}
}
