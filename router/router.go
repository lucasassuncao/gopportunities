package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router using GIN's Default Configuration
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8090")
}
