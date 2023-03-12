// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Let's model the users table
type Equipment struct {
	EquipmentID    int64
	Email     string
	CreatedAt time.Time
}

// Setup dependency injection
type EquipmentModel struct {
	DB *sql.DB
}

// Write SQL code to access the database
// TODO
// Creating a Get Method for Users table
func (m *EquipmentModel) Get() (*Equipment, error) {
	var q Equipment

	statement := `
	            SELECT equipment_id, body
				FROM equipments
				ORDER BY RANDOM()
				LIMIT 1
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.EquipmentID, &q.Email)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *EquipmentModel) Insert(body string) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO equipments(body)
				VALUES($1)
				RETURNING equipment_id				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}