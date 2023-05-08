// Filename: internal/models/logs.go

package models

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Log represents the log data structure
type Log struct {
	ID       int64     // Log ID
	Username string    // Username of the user who signed in
	DateTime time.Time // Date and time of the sign-in event
}

// Setup dependency injection
type LogModel struct {
	DB *sql.DB
}

// // InsertLog inserts a log entry into the PostgreSQL database
// func InsertLog(username string) error {
// 	// Connect to PostgreSQL database
// 	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Get the current date and time
// 	currentTime := time.Now()

// 	// Insert log entry into the database
// 	_, err = db.Exec("INSERT INTO logs (username, datetime) VALUES ($1, $2)", username, currentTime)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Log has been inserted successfully!")
// 	return nil
// }

// Create a Log
func (m *LogModel) Insert(username string) (int64, error) {
	// id will be used to stored the unique identifier returned by
	// PostgreSQL after adding the row to the table
	var id int64
	datetime := time.Now()

	statement :=
		`
		        INSERT INTO logs (username, datetime)
				VALUES ($1, $2)
				RETURNING log_id
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(id, username, datetime)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Read a particular Log
func (m *LogModel) Get() (*Log, error) {
	var q Log

	statement := `
	            SELECT id, username, datetime  
				FROM logs
				WHERE id = $1
				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.ID, &q.Username, &q.DateTime)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
