package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hariszaki17/library-management/book-service/cache"
	"github.com/hariszaki17/library-management/book-service/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetBooks(ctx context.Context, page, limit int) ([]*models.Book, error)
}

type bookRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewBookRepository(db *gorm.DB, cache *cache.Cache) BookRepository {
	return &bookRepository{db: db, cache: cache}
}

func (r *bookRepository) GetBooks(ctx context.Context, page, limit int) ([]*models.Book, error) {
	cacheKey := fmt.Sprintf("GetBooks, page:%d, limit:%d", page, limit)
	cachedBook, err := r.cache.Get(cacheKey)
	if err == nil && cachedBook != "" {
		// Deserialize cachedBook and return
		books := []*models.Book{}
		err := json.Unmarshal([]byte(cachedBook), &books)
		if err != nil {
			return nil, err
		}
		return books, nil
	}

	// Calculate the correct offset
	offset := (page - 1) * limit

	var books []*models.Book
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		return nil, err
	}

	// Cache the result
	serializedBooks, err := json.Marshal(books)
	if err != nil {
		return nil, err
	}
	err = r.cache.Set(cacheKey, string(serializedBooks), time.Minute)
	if err != nil {
		log.Printf("Failed to set cache: %v", err)
	}
	return books, nil
}
