package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)


func Init() error {
	var err error

	db, err = initializationSQLite()

	if err != nil {
		return fmt.Errorf("error initialization SQLite: %v",err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger =  NewLogger(prefix)

	return logger
}
