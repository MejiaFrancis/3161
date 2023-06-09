// Filename: internal/models/reservation.go
// this file is used to show the fields of a reservation
package models

import (
	"context"
	"database/sql"
	"time"
)

// The Question model will represent a single question in our questions table
type login struct {
	LoginID   int64
	UserID    int64
	Email     string
	Password  string
	CreatedAt time.Time
}

// The QuestionModel type will encapsulate the
// DB connection pool that will be initialized
// in the main() function
type LoginModel struct {
	DB *sql.DB
}

// The Insert() function stores a question into the  table
func (m *LoginModel) Insert(Email string, Password string) (int64, error) {
	// id will be used to stored the unique identifier returned by
	// PostgreSQL after adding the row to the table
	var id int64
	statement :=
		` Select qurey to find data on user
						
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *LoginModel) Get() (*login, error) {
	var Login login

	statement :=
		`
							SELECT reservations_id, user_id, people_count
							FROM reservations
							ORDER BY RANDOM()
							LIMIT 1
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&Login.Email, &Login.Password)
	if err != nil {
		return nil, err
	}
	return &Login, nil
}
