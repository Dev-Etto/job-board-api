package main

import (
	"github.com/Dev-Etto/job-board-api/config"
	"github.com/Dev-Etto/job-board-api/router"
)


var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
