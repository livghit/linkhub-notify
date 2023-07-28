package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create the connection to the database
func ConnectDatabase(databasePath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("failed the db connection")
	}

	return db
}
