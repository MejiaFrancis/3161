// Filename: internal/models/roles.go
package models

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Let's model the users table
type Role struct {
	RoleID            int64
	RoleStudnet       string
	RoleAdministrator string
	RoleTeacher       string
}

// // Role represents the role of a user
// type Role string

// const (
// 	RoleStudent       Role = "Student"
// 	RoleAdministrator Role = "Administrator"
// 	RoleTeacher       Role = "Teacher"
// )

// Setup dependency injection
type RoleModel struct {
	DB *sql.DB
}

// func main() {
// 	// Prompt user to select a role
// 	fmt.Println("Select a role: ")
// 	fmt.Println("1. Student")
// 	fmt.Println("2. Administrator")
// 	fmt.Println("3. Teacher")

// 	fmt.Scanln(&roleInput)

// 	// Map user input to Role type

// 	// Connect to PostgreSQL database
// 	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Insert role into the database
// 	_, err = db.Exec()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// Creating a Get Method for Roles table
func (m *RoleModel) Get() (*Role, error) {
	var q Role

	statement := `
	            SELECT rolestudent, roleadministrator, roleteacher 
				FROM roles
				ORDER BY RANDOM()
				LIMIT 1
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.RoleStudnet, &q.RoleAdministrator, &q.RoleTeacher)
	if err != nil {
		return nil, err
	}
	return &q, err
}

func (m *RoleModel) Insert(rolestudent, roleadministrator, roleteacher string) (int64, error) {
	//var role Role
	//var roleInput int
	var id int64

	query := `
	        INSERT INTO roles (rolestudent, roleadministrator, roleteacher) 
			VALUES ($1, $2, $3)
			returning roleid
	        `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query).Scan(id)
	if err != nil {
		return 0, nil
	}
	return id, nil
}

// switch roleInput {
// case 1:
// 	role = "RoleStudent"
// case 2:
// 	role = "RoleAdministrator"
// case 3:
// 	role = "RoleTeacher"
// default:
// 	log.Fatal("Invalid role selected")
// }
// fmt.Println("Role has been inserted successfully!")
