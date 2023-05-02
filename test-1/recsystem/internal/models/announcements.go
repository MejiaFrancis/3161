// Filename: internal/models/announcements.go

package models

import (
	"context"
	"database/sql"
	//"go/constant"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Announcement represents the announcement data structure
type Announcement struct {
	ID      int    // Announcement ID
	Subject string // Announcement subject
	Content string // Announcement content
}

// Setup dependency injection
type AnnouncementModel struct {
	DB *sql.DB
}


// Creating a Get Method for Announcements table
func (m *AnnouncementModel) Get() (*Announcement, error) {
	var q Announcement

	statement := `
	            SELECT id, subject, content
				FROM Announcements
				ORDER BY RANDOM()
				LIMIT 1
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.ID, &q.Subject, &q.Content)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post Announcements entered into the database
func (m *AnnouncementModel) Insert(subject string, content string) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO Announcements(subject, content)
				VALUES($1)
				RETURNING Announcementid				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, subject, content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
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
