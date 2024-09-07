package models

import (
	"time"

	"gorm.io/gorm"
)

type BorrowingRecord struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	BookID     uint       `json:"book_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`
}

type BorrowingCount struct {
	BookID uint `json:"book_id"`
	Count  int  `json:"count"`
}
