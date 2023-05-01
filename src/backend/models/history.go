package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type History struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserQuery string    `json:"user_query"`
	Response  string    `json:"response"`
	SessionId uuid.UUID `json:"session_id" gorm:"type:uuid;index"`
}
