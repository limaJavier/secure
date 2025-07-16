package persistence

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDb(inMemory bool) (*gorm.DB, error) {
	//** Establish database connection
	var db *gorm.DB
	var err error
	if inMemory {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	} else {
		dbDir := filepath.Join(xdg.DataHome, "secure")
		dbPath := filepath.Join(dbDir, "secure.db")

		// Ensure the directory exists
		if err := os.MkdirAll(dbDir, 0700); err != nil {
			return nil, fmt.Errorf("cannot create db directory: %v", err)
		}

		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	//** Ensure database schema is correct
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
