// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Create a user
var (
	ErrNoRecord           = errors.New("no matching record found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail     = errors.New("duplicate email")
	ErrEditConflict       = errors.New("edit conflict")
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
	Password     []byte
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
	            SELECT userid, email, first_name, last_name, age, address, phone_number, roles_id, password_hash, status
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
func (m *UserModel) Insert(email string, first_name string, password string) error {
	// let's hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `
	            INSERT INTO users(first_name, email,  password_hash)
				VALUES($1, $2, $3,)
			             `

	// Collect the data fields into a slice
	// args := []interface{}{
	// 	user.Email, user.First_name, user.Last_name,
	// 	user.Age, user.Address, user.Phone_number,
	// 	user.Roles_id,
	// 	//user.Password,
	// 	user.Status,
	// }
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = m.DB.ExecContext(ctx, query, first_name, email, hashedPassword)
	if err != nil {
		switch {
		case err.Error() == `pgx: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	// check if there is a row in the table foer the email provided
	query := `
	         SELECT id, password_hash
			 FROM users
			 WHERE email = $1
	         `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	} ///handling error
	// the user does exist
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	// password is correct
	return id, nil
}

func (m UserModel) Update(user *User) error {
	// Create a query
	query := `
		UPDATE user
		SET userid = $1, first_name = $2, last_name = $3,
		    age = $4, address = $5, phone_number = $6,
			email = $7, roles_id = $8, password_hash = $9
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
		return ErrNoRecord
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
		return ErrNoRecord
	}
	return nil
}
