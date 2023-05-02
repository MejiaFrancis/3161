// Filename: internal/models/models.go

package models

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// A wrapper for our data models
type Models struct {
	Users             UserModel
	Equipments        EquipmentModel
	Roles             RoleModel
	Reservations      ReservationModel
	Logs              LogModel
	Feedbacks         FeedbackModel
	EquipmentUsageLog EquipmentUsageLogModel
	EquipmentTypes    EquipmentTypeModel
	Announcements     AnnouncementModel
}

// NewModels() allows us to create a new Models
func NewModels(db *sql.DB) Models {
	return Models{
		Users:             UserModel{DB: db},
		Equipments:        EquipmentModel{DB: db},
		Roles:             RoleModel{DB: db},
		Reservations:      ReservationModel{DB: db},
		Logs:              LogModel{DB: db},
		Feedbacks:         FeedbackModel{DB: db},
		EquipmentUsageLog: EquipmentUsageLogModel{DB: db},
		EquipmentTypes:    EquipmentTypeModel{DB: db},
		Announcements:     AnnouncementModel{DB: db},
	}
}
