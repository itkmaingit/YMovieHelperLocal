package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itkmaingit/YMovieHelper/controllers/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	// router.POST("/login", handlers.Login)
	router.GET("health_check", handlers.HealthCheck)
	// router.GET("/auth", handlers.Auth)
	// router.POST("/user", handlers.CreateUser)
	// user := router.Group("/user")
	// user.Use(middlewares.AccessControllMiddleware())
	// {
	router.GET("/softwares", handlers.GetSoftwaresAndProjects)
	router.POST("/softwares", handlers.CreateSoftware)

	software := router.Group("/softwares/:softwareID")
	{
		software.PUT("", handlers.UpdateSoftware)
		software.DELETE("", handlers.DeleteSoftware)

		software.GET("/characters", handlers.GetCharacters)
		software.POST("/characters", handlers.CreateCharacters)

		software.POST("/projects", handlers.CreateProject)

		character := software.Group("/characters/:characterID")
		{
			character.PUT("", handlers.UpdateCharacter)
			character.DELETE("", handlers.DeleteCharacter)

			emotion := character.Group("/emotions")
			{
				emotion.GET("", handlers.GetEmotions)
				emotion.POST("", handlers.CreateCharacterEmotions)
				emotion.PUT("", handlers.UpdateEmotion)
				emotion.DELETE("", handlers.DeleteEmotion)
			}
		}

		project := software.Group("/projects/:projectID")
		{
			project.PUT("", handlers.UpdateProject)
			project.DELETE("", handlers.DeleteProject)

			items := project.Group("/items")
			{
				items.GET("", handlers.GetItems)
				items.POST("/resolve_single_item", handlers.ResolveSingleItem)
				items.POST("/create_single_item", handlers.UploadSingleItem)
				items.POST("/upload_multiple_item", handlers.UploadMultipleItem)
				items.POST("/upload_dynamic_item", handlers.UploadDynamicItem)

				items.PUT("/single_item", handlers.UpdateSingleItem)
				items.PUT("/multiple_item", handlers.UpdateMultipleItem)
				items.PUT("/dynamic_item", handlers.UpdateDynamicItem)

				items.DELETE("/single_item", handlers.DeleteSingleItem)
				items.DELETE("/multiple_item", handlers.DeleteMultipleItem)
				items.DELETE("/dynamic_item", handlers.DeleteDynamicItem)
			}

			rules := project.Group("/rules")
			{
				rules.GET("", handlers.GetRule)
				rules.POST("", handlers.CreateRule)
			}

			makeYMMP := project.Group("/make_ymmp")
			{
				makeYMMP.GET("/can_make_ymmp", handlers.CheckRules)
				makeYMMP.POST("/resolve_scenario", handlers.ResolveScenario)
				makeYMMP.POST("/resolve_ymmp", handlers.MakeYMMP)
			}

		}

	}

	// }
	download := router.Group("/download/:projectID")
	{
		download.GET("scenario_csv", handlers.DownloadScenarioCSV)
		download.GET("scenario_txt", handlers.DownloadScenarioTXT)
		download.GET("complete_ymmp", handlers.DownloadCompleteYMMP)
	}

}
