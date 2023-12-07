package main

import (
	"github.com/lucasassuncao/gopportunities/config"
	"github.com/lucasassuncao/gopportunities/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	// Load configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config error: %v", err.Error())
		return
	}

	// Initialize Router
	router.Initialize()
}
