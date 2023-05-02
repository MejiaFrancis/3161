// Filename: internals/models/reservations.go

package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Reservation represents the reservation data structure
type Reservation struct {
	ID              int       // Reservation ID
	UserID          int       // User ID
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
// InsertReservation inserts a reservation into the PostgreSQL database
func InsertReservation(userID int, reservationDate, reservationTime string, duration, peopleCount int, notes string) error {
	// Connect to PostgreSQL database
	db, err := sql.Open("PostgreSQL DSN", "RECSYSTEM_DB_DSN")
	if err != nil {
		return err
	}
	defer db.Close()

	// Get the current timestamp for reservation creation
	createdAt := time.Now()

	// Insert reservation into the database with approval status set to false (pending)
	_, err = db.Exec("INSERT INTO reservations (user_id, reservation_date, reservation_time, duration, people_count, notes, approval, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		userID, reservationDate, reservationTime, duration, peopleCount, notes, false, createdAt)
	if err != nil {
		return err
	}

	log.Println("Reservation has been created and pending approval.")
	return nil
}
