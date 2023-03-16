// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Let's model the users table
type Equipment struct {
	EquipmentID       int64
	Name              string
	Image             string
	Equipment_type_id int
	Status            bool
	CreatedAt         time.Time
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
	            SELECT equipmentid, name, image, equipment_type_id, status
				FROM equipments
				ORDER BY RANDOM()
				LIMIT 1
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.EquipmentID, &q.Name, &q.Image, &q.Equipment_type_id, &q.Status)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *EquipmentModel) Insert(name string, image string, equipment_type_id int, status bool) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO equipments(name, image, equipment_type_id, status)
				VALUES($1)
				RETURNING equipmentid				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, name, image, equipment_type_id, status).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
