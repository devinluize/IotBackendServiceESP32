package database

import (
	"gorm.io/gorm"
	"log"
)

// DropAllDatabase drops all tables in the specified database.
// It retrieves the list of all table names and iteratively drops each table.
// Logs success and failure for each table drop operation.
func DropAllDatabase(db *gorm.DB) error {
	// Get the list of all tables
	var tableNames []string
	db.Raw("SELECT table_name FROM article_schema.tables WHERE table_type = 'BASE TABLE'").Scan(&tableNames)

	// Drop each table
	for _, tableName := range tableNames {
		err := db.Migrator().DropTable(tableName)
		if err != nil {
			log.Printf("Failed to drop table %s: %v", tableName, err)
			return err
		} else {
			log.Printf("Successfully dropped table %s", tableName)
		}
	}
	return nil
}
