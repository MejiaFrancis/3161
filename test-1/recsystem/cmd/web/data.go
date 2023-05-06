// Filename: cmd/web/data.go
package main

import "gibhub.com/MejiaFrancis/3161/3162/test-1/recsystem/internal/models"

type templateData struct {
	//Question *models.Question
	User *models.User
	//Role              *models.RoleModel
	//Reservations      *models.ReservationModel
	//Logs              *models.LogModel
	//Feedback          *models.FeedbackModel
	//Equipmentusagelog *models.EquipmentUsageLogMode
	//Equipment_types   *models.EquipmentTypeModel
	//Announcements     *models.AnnouncementModel
	Flash           string
	CSRFToken       string
	IsAuthenticated bool
}
