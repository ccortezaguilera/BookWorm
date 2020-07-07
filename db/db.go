package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
)

// GetDBPath : Gives the file path of the database file.
func GetDBPath(dbName string) string {
	var cwd, _ = os.Getwd()
	return filepath.Join(cwd, dbName)
}

// Migrate : Migrates the structs from the array of given structs.
func Migrate(db *gorm.DB, models []struct{}) {
	for _, model := range models {
		db.AutoMigrate(&model)
	}
}

// CreateDB : creates the db file if it doesn't exist.
func CreateDB(dbName string) {
	dbPath := GetDBPath(dbName)
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("Creating %s Database", dbName)
		file, creationErr := os.Create(dbPath)
		if creationErr != nil {
			log.Fatal(creationErr.Error())
		}
		file.Close()
		log.Printf("Created %s Database", dbName)
	}
}
