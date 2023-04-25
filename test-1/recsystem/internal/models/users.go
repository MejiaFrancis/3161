// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"errors"
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
func (m *UserModel) Get(userid int64) (*User, error) {
	var q User

	statement := `
	            SELECT userid, email, first_name, last_name, age, address, phone_number, roles_id, password, status
				FROM users
				ORDER BY RANDOM()
				LIMIT 1
	             `

	// Declare a user variable to hold the returned data
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Execute the query using QueryRow()
	err := m.DB.QueryRowContext(ctx, statement, userid).Scan(
		&user.UserID,
		&user.CreatedAt,
		&user.First_name,
		&user.Last_name,
		&user.Age,
		&user.Address,
		&user.Phone_number,
		&user.Email,
		&user.Roles_id,
		&user.Password,
		&user.Status,
	)

	err = m.DB.QueryRowContext(ctx, statement).Scan(&q.UserID, &q.Email, &q.First_name, &q.Last_name, &q.Age, &q.Address, &q.Phone_number, &q.Roles_id, &q.Password, &q.Status)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *UserModel) Insert(user *User) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO users(email, first_name, last_name, age, address, phone_number, roles_id, password, status)
				VALUES($1)
				RETURNING userid				
	             `

	// Collect the data fields into a slice
	args := []interface{}{
		user.Email, user.First_name, user.Last_name,
		user.Age, user.Address, user.Phone_number,
		user.Roles_id,
		user.Password, user.Status,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, args...).Scan(&user.UserID)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m UserModel) Update(user *User) error {
	// Create a query
	query := `
		UPDATE user
		SET userid = $1, first_name = $2, last_name = $3,
		    age = $4, address = $5, phone_number = $6,
			email = $7, roles_id = $8, password = $9
		WHERE status = $10
		RETURNING status
	`
	args := []interface{}{
		user.UserID,
		user.CreatedAt,
		user.First_name,
		user.Last_name,
		user.Age,
		user.Address,
		user.Phone_number,
		user.Email,
		user.Roles_id,
		user.Password,
		user.Status,
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Check for edit conflicts
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Status)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):

			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

// Delete() removes a specific User
func (m UserModel) Delete(id int64) error {
	// Ensure that there is a valid id
	if id < 1 {
		return ErrRecordNotFound
	}
	// Create the delete query
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query
	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	// Check how many rows were affected by the delete operation. We
	// call the RowsAffected() method on the result variable
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Check if no rows were affected
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
