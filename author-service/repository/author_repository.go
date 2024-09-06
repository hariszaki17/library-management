package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hariszaki17/library-management/author-service/cache"
	"github.com/hariszaki17/library-management/author-service/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAuthors(ctx context.Context, page, limit int) ([]*models.Author, error)
}

type authorRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewAuthorRepository(db *gorm.DB, cache *cache.Cache) AuthorRepository {
	return &authorRepository{db: db, cache: cache}
}

func (r *authorRepository) GetAuthors(ctx context.Context, page, limit int) ([]*models.Author, error) {
	cacheKey := fmt.Sprintf("GetAuthors, page:%d, limit:%d", page, limit)
	cachedAuthor, err := r.cache.Get(cacheKey)
	if err == nil && cachedAuthor != "" {
		// Deserialize cachedAuthor and return
		books := []*models.Author{}
		err := json.Unmarshal([]byte(cachedAuthor), &books)
		if err != nil {
			return nil, err
		}
		return books, nil
	}

	// Calculate the correct offset
	offset := (page - 1) * limit

	var books []*models.Author
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		return nil, err
	}

	// Cache the result
	serializedAuthors, err := json.Marshal(books)
	if err != nil {
		return nil, err
	}
	err = r.cache.Set(cacheKey, string(serializedAuthors), time.Minute)
	if err != nil {
		log.Printf("Failed to set cache: %v", err)
	}
	return books, nil
}
