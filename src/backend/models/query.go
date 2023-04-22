package models

import "gorm.io/gorm"

type Query struct {
	gorm.Model
	Query    string `json:"query"`
	Response string `json:"response"`
}
