// Filename: internals/models/reservations.go

package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Reservation represents the reservation data structure
type Reservation struct {
	ID              int64     // Reservation ID
	UserID          int64     // User ID
	ReservationDate string    // Reservation date
	ReservationTime string    // Reservation time
	Duration        int       // Duration in minutes
	PeopleCount     int       // Number of people
	Notes           string    // Additional notes
	Approval        bool      // Approval status (true if approved, false if pending)
	CreatedAt       time.Time // Timestamp of reservation creation
}

// Setup dependency injection
type ReservationModel struct {
	DB *sql.DB
}

// // InsertReservation inserts a reservation into the PostgreSQL database
// func InsertReservation(userID int, reservationDate, reservationTime string, duration, peopleCount int, notes string) error {
// 	// Connect to PostgreSQL database
// 	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Get the current timestamp for reservation creation
// 	createdAt := time.Now()

// 	// Insert reservation into the database with approval status set to false (pending)
// 	_, err = db.Exec("INSERT INTO reservations (user_id, reservation_date, reservation_time, duration, people_count, notes, approval, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
// 		userID, reservationDate, reservationTime, duration, peopleCount, notes, false, createdAt)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Reservation has been created and pending approval.")
// 	return nil
// }

// Create Reservation
func (m *ReservationModel) Insert(userID int64, reservationDate, reservationTime string, duration, peopleCount int, notes string) (int64, error) {
	// id will be used to stored the unique identifier returned by
	// PostgreSQL after adding the row to the table
	var id int64
	// Get the current timestamp for reservation creation
	createdAt := time.Now()

	statement :=
		`
				INSERT INTO reservation(user_id, reservation_date, reservation_time, duration, people_count, notes, approval, created_at)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING reservation_id
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(id, userID, reservationDate, reservationTime, duration, peopleCount, notes, false, createdAt)
	if err != nil {
		return 0, err
	}
	log.Println("Reservation has been created and pending approval.")
	return id, nil
}

// Read (Get) Reservation
func (m *ReservationModel) Get() (*Reservation, error) {
	var q Reservation

	statement :=
		`
				SELECT reservation_id, user_id, reservation_date, reservation_time, duration, people_count, notes, approval, created_at
			    FROM reservation
				ORDER BY RANDOM()
				LIMIT 1
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.ID, &q.UserID, &q.ReservationDate, &q.ReservationTime, &q.Duration, &q.PeopleCount, &q.Notes, &q.Approval, &q.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
