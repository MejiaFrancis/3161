// Filename: internal/models/logs.go

package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Log represents the log data structure
type Log struct {
	ID       int       // Log ID
	Username string    // Username of the user who signed in
	DateTime time.Time // Date and time of the sign-in event
}

// InsertLog inserts a log entry into the PostgreSQL database
func InsertLog(username string) error {
	// Connect to PostgreSQL database
	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
	if err != nil {
		return err
	}
	defer db.Close()

	// Get the current date and time
	currentTime := time.Now()

	// Insert log entry into the database
	_, err = db.Exec("INSERT INTO logs (username, datetime) VALUES ($1, $2)", username, currentTime)
	if err != nil {
		return err
	}

	log.Println("Log has been inserted successfully!")
	return nil
}
