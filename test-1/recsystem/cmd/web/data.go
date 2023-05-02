// Filename: cmd/web/data.go
package main

import (
	"github.com/MejiaFrancis/3161/3162/test-1/recsystem/internal/models"
)

type templateData struct {
	Question *models.Question
	User     *models.User
	Flash    string
}
