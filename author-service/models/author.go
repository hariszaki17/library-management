package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
	BirthDate time.Time `json:"birth_date"`
}
