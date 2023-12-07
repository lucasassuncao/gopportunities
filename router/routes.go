package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasassuncao/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/opening", handler.CreateOpeningHandler)   // Create
		v1.GET("/opening", handler.ShowOpeningHandler)      // Read
		v1.PUT("/opening", handler.UpdateOpeningHandler)    // Update
		v1.DELETE("/opening", handler.DeleteOpeningHandler) // Delete
		v1.GET("/openings", handler.ListOpeningsHandler)    // List
	}
}
