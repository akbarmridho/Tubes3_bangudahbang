package models

import "time"

type Query struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Query     string    `json:"query"`
	Response  string    `json:"response"`
}
