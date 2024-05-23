package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RegisterMiddleware(router *gin.Engine) {
	router.Use(CORSMiddleware(router))
	router.Use(InitializeSessionStore())
}
