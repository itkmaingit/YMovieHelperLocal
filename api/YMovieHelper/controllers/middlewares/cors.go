package middlewares

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(r *gin.Engine) gin.HandlerFunc {

	frontendEndpoint := os.Getenv("FrontendEndpoint")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frontendEndpoint}                          // 許可するオリジンを指定。ここでは例としてローカルのポート3000を許可
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}            // 許可するHTTPメソッドを指定
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // 許可するリクエストヘッダーを指定
	config.AllowCredentials = true                                            // 資格情報（クッキーなど）を含むリクエストを許可するかどうか

	return cors.New(config)
}
