package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UserQuery string `json:"user_query"`
	Response  string `json:"response"`
}
