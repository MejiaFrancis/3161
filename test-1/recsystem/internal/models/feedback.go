// Filename: internal/models/feedback.go

package models

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Feedback represents the feedback data structure
type Feedback struct {
	ID      int    // Feedback ID
	Name    string // Name of the user leaving feedback
	Message string // Feedback message
}

// Setup dependency injection
type FeedbackModel struct {
	DB *sql.DB
}

// // InsertFeedback inserts feedback into the PostgreSQL database
// func InsertFeedback(feedback Feedback) error {
// 	// Connect to PostgreSQL database
// 	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Insert feedback into the database
// 	_, err = db.Exec("INSERT INTO feedback (name, message) VALUES ($1, $2)", feedback.Name, feedback.Message)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Feedback has been inserted successfully!")
// 	return nil
// }

// Create Feedback
func (m *FeedbackModel) Insert(name, message string) (int64, error) {
	// id will be used to stored the unique identifier returned by
	// PostgreSQL after adding the row to the table
	var id int64

	statement :=
		`
		        INSERT INTO feedback (name, message) 
				VALUES ($1, $2)
				RETURNING feedback_id
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(id, name, message)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Read (View) Feedback
func (m *FeedbackModel) Get() (*Feedback, error) {
	var q Feedback

	statement := `
	            SELECT id, name, message 
				FROM feedback
				WHERE id = $1
				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.ID, &q.Name, &q.Message)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
