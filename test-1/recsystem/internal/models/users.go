// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Let's model the users table
type User struct {
	UserID       int64
	Email        string
	First_name   string
	Last_name    string
	Age          int
	Address      int32
	Phone_number int16
	Roles_id     int
	Password     string
	Status       bool
	CreatedAt    time.Time
}

// Setup dependency injection
type UserModel struct {
	DB *sql.DB
}

// Write SQL code to access the database
// TODO
// Creating a Get Method for Users table
func (m *UserModel) Get() (*User, error) {
	var q User

	statement := `
	            SELECT userid, email, first_name, last_name, age, address, phone_number, roles_id, password, status
				FROM users
				ORDER BY RANDOM()
				LIMIT 1
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.UserID, &q.Email, &q.First_name, &q.Last_name, &q.Age, &q.Address, &q.Phone_number, &q.Roles_id, &q.Password, &q.Status)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *UserModel) Insert(email string, first_name string, last_name string, age int, address int32, phone_number int16, roles_id int, password string, status bool) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO users(email, first_name, last_name, age, address, phone_number, roles_id, password, status)
				VALUES($1)
				RETURNING userid				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, email, first_name, last_name, age, address, phone_number, roles_id, password, status).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
