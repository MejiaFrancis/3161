// Filename: internal/models/announcements.go

package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Announcement represents the announcement data structure
type Announcement struct {
	ID      int    // Announcement ID
	Subject string // Announcement subject
	Content string // Announcement content
}

// InsertAnnouncement inserts an announcement into the PostgreSQL database
func InsertAnnouncement(subject, content string) error {
	// Connect to PostgreSQL database
	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert announcement into the database
	_, err = db.Exec("INSERT INTO announcements (subject, content) VALUES ($1, $2)", subject, content)
	if err != nil {
		return err
	}

	log.Println("Announcement has been created.")
	return nil
}
