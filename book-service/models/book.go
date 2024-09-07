package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `json:"title"`
	AuthorID    uint      `json:"author_id"`
	CategoryID  uint      `json:"category_id"`
	ISBN        string    `json:"isbn" gorm:"unique"`
	PublishedAt time.Time `json:"published_at"`
	Stock       uint      `json:"stock"`
}

type BookRecommendation struct {
	*Book
	BorrowedCount uint `json:"borrowed_count"`
}
