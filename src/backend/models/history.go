package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	UserQuery string    `json:"user_query"`
	Response  string    `json:"response"`
	SessionId uuid.UUID `json:"session_id" gorm:"type:uuid;index"`
}
