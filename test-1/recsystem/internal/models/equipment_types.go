// Filename: internal/models/equipment_types.go

package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// EquipmentType represents the equipment type data structure
type EquipmentType struct {
	ID       int    // Equipment type ID
	TypeName string // Equipment type name
}

// InsertEquipmentType inserts an equipment type into the PostgreSQL database
func InsertEquipmentType(typeName string) error {
	// Connect to PostgreSQL database
	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert equipment type into the database
	_, err = db.Exec("INSERT INTO equipment_types (type_name) VALUES ($1)", typeName)
	if err != nil {
		return err
	}

	log.Println("Equipment type has been created.")
	return nil
}
