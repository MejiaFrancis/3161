// Filename: internal/models/roles.go
package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Role represents the role of a user
type Role string

const (
	RoleStudent       Role = "Student"
	RoleAdministrator Role = "Administrator"
	RoleTeacher       Role = "Teacher"
)

func main() {
	// Prompt user to select a role
	fmt.Println("Select a role: ")
	fmt.Println("1. Student")
	fmt.Println("2. Administrator")
	fmt.Println("3. Teacher")
	var roleInput int
	fmt.Scanln(&roleInput)

	// Map user input to Role type
	var role Role
	switch roleInput {
	case 1:
		role = RoleStudent
	case 2:
		role = RoleAdministrator
	case 3:
		role = RoleTeacher
	default:
		log.Fatal("Invalid role selected")
	}

	// Connect to PostgreSQL database
	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert role into the database
	_, err = db.Exec("INSERT INTO roles (role) VALUES ($1)", role)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Role has been inserted successfully!")
}
