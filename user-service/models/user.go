package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type AuthUseCaseResp struct {
	User  `json:"user"`
	Token string `json:"token"`
}
