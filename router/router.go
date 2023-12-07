package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router using GIN's Default Configuration
	router := gin.Default()

	// Defining Route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8090") // listen and serve on 0.0.0.0:8080
}
