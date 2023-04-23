// Filename: internal/models/feedback.go

package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Feedback represents the feedback data structure
type Feedback struct {
	ID      int    // Feedback ID
	Name    string // Name of the user leaving feedback
	Message string // Feedback message
}

// InsertFeedback inserts feedback into the PostgreSQL database
func InsertFeedback(feedback Feedback) error {
	// Connect to PostgreSQL database
	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert feedback into the database
	_, err = db.Exec("INSERT INTO feedback (name, message) VALUES ($1, $2)", feedback.Name, feedback.Message)
	if err != nil {
		return err
	}

	log.Println("Feedback has been inserted successfully!")
	return nil
}
