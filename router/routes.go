package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasassuncao/gopportunities/handler"

	docs "github.com/lucasassuncao/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.POST("/opening", handler.CreateOpeningHandler)   // Create
		v1.GET("/opening", handler.ShowOpeningHandler)      // Read
		v1.PUT("/opening", handler.UpdateOpeningHandler)    // Update
		v1.DELETE("/opening", handler.DeleteOpeningHandler) // Delete
		v1.GET("/openings", handler.ListOpeningHandler)     // List
	}

	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
