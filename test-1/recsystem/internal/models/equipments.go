// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Let's model the users table
type Equipment struct {
	EquipmentID       int64
	Name              string
	Image             string
	Equipment_type_id int
	Status            bool
	Availability      bool
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
func (m *EquipmentModel) Insert(name string, image string, equipment_type_id int, status bool, availability bool) (int64, error) {

	statement := `
	INSERT INTO equipment (name, image, equipment_type_id, status, availability)
	VALUES($1,$2,$3,$4,$5)
	RETURNING equipment_id				
	`
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	st, err := m.DB.Prepare(statement)
	if err != nil {
		return 0, err
	}
	defer st.Close()

	result, err := st.Exec(name, image, equipment_type_id, status, availability)
	// err := m.DB.QueryRowContext(ctx, statement, name /*image,*/, equipment_type_id, status, availability).Scan(&id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 1 {
		fmt.Println("Insertion successful")
	} else {
		fmt.Println("Insertion failed")
	}

	return rowsAffected, nil
}

// function to delete a piece of equipment from the database
func (m *EquipmentModel) Delete(equip_id int64) (int64, error) {
	var id int64

	statement := `
	DELETE FROM equipment
	WHERE equipment_id = ($1)				
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// does the deletion and returns the id of the deleted equipment to be used for
	// confirmation of deletion
	err := m.DB.QueryRowContext(ctx, statement, equip_id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// function to mark a peice of equipment as borrowed in the database
func (m *EquipmentModel) Borrow(equip_id int64) (int64, error) {
	var id int64

	statement := `
	UPDATE equipment
	SET status = false
	WHERE equipment_id = ($1);		
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, statement, equip_id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// function to mark a peice of equipment as available for borrowing in the database
func (m *EquipmentModel) Return(equip_id int64) (int64, error) {
	var id int64

	statement := `
	UPDATE equipment
	SET status = true
	WHERE equipment_id = ($1);		
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, statement, equip_id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
