package middlewares

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitializeSessionStore() gin.HandlerFunc {

	sessionSecretKey := os.Getenv("SessionSecretKey")
	store := cookie.NewStore([]byte(sessionSecretKey))
	return sessions.Sessions("ymovieHelper", store)
}
