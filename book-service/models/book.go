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
	ISBN        string    `json:"isbn"`
	PublishedAt time.Time `json:"published_at"`
	Stock       uint      `json:"stock"`
}
