package persistence

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDb(inMemory bool) (*gorm.DB, error) {
	// Establish database connection
	var db *gorm.DB
	var err error
	if inMemory {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	} else {
		db, err = gorm.Open(sqlite.Open("secure.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		err = os.Chmod("secure.db", 0666) // Allow read/write to all users
		if err != nil {
			return nil, err
		}
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
