package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/controllers/middlewares"
	"github.com/itkmaingit/YMovieHelper/controllers/routes"
	"github.com/itkmaingit/YMovieHelper/repository"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			YMovieHelper API Server
// @version		1.0
// @license.name	dango
// @description	This is a YMovieHelper API server.
// @termsOfService	http://swagger.io/terms/
// @contact.email	itkmain.git@gmail.com
// @host			localhost:8080
// @BasePath		/
func main() {
	// 環境変数の読み込み
	_ = godotenv.Load()

	router := gin.New()

	// Apply Recovery middleware
	router.Use(gin.Recovery())

	// Apply customized Logger middleware
	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path != "/health_check" {
			gin.Logger()(c)
		}
	})

	router.SetTrustedProxies([]string{os.Getenv("TrustedProxyIPAddress")})

	// データベース接続の初期化
	db := repository.GetDB()
	defer db.Close()

	// Settings
	middlewares.RegisterMiddleware(router)
	routes.RegisterRoutes(router)
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run
	router.Run(":80")

}
