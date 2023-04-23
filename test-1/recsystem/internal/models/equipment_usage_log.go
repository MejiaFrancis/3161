// Filename: internal/models/equipment_usage_log.go

package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// EquipmentUsageLog represents the equipment usage log data structure
type EquipmentUsageLog struct {
	ID             int       // Log ID
	EquipmentID    int       // Equipment ID
	UserID         int       // User ID
	LogID          int       // Log ID (e.g. log entry ID)
	TimeBorrowed   time.Time // Time when the equipment was borrowed
	TimeReturned   time.Time // Time when the equipment was returned
	ReturnedStatus bool      // Status of equipment return (true if returned, false if not)
}

// InsertEquipmentUsageLog inserts an equipment usage log entry into the PostgreSQL database
func InsertEquipmentUsageLog(equipmentID, userID int, logID int, timeBorrowed, timeReturned time.Time, returnedStatus bool) error {
	// Connect to PostgreSQL database
	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert equipment usage log entry into the database
	_, err = db.Exec("INSERT INTO equipment_usage_logs (equipment_id, user_id, log_id, time_borrowed, time_returned, returned_status) VALUES ($1, $2, $3, $4, $5, $6)",
		equipmentID, userID, logID, timeBorrowed, timeReturned, returnedStatus)
	if err != nil {
		return err
	}

	log.Println("Equipment usage log has been inserted successfully!")
	return nil
}
