package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	// return errors.New("FAKE ERROR")

	var err error
	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return err
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Initialize Logger
	logger = NewLogger(prefix)
	return logger
}
