package config

import (
	"os"

	"github.com/Dev-Etto/job-board-api/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializationSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	_ , err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schema.Opening{})

	if err != nil {
		logger.Errorf("sqlite automigrate error: %v", err)
	}

	return db, err
}
