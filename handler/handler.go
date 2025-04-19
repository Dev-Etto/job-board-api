package handler

import (
	"github.com/Dev-Etto/job-board-api/config"
	"gorm.io/gorm"
)

var  (
	logger *config.Logger
	db *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
