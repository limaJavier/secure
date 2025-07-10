package persistence

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDb() (*gorm.DB, error) {
	// Establish database connection
	db, err := gorm.Open(sqlite.Open("secure.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Ensure database schema is correct
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&Entry{})
	if err != nil {
		return nil, err
	}

	return db, err
}
