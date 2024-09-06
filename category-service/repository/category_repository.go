package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hariszaki17/library-management/category-service/models"
	"github.com/hariszaki17/library-management/proto/cache"

	// "github.com/hariszaki17/library-management/category-service/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context, page, limit int) ([]*models.Category, error)
}

type categoryRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewCategoryRepository(db *gorm.DB, cache *cache.Cache) CategoryRepository {
	return &categoryRepository{db: db, cache: cache}
}

func (r *categoryRepository) GetCategories(ctx context.Context, page, limit int) ([]*models.Category, error) {
	cacheKey := fmt.Sprintf("GetCategories, page:%d, limit:%d", page, limit)
	cachedBook, err := r.cache.Get(cacheKey)
	if err == nil && cachedBook != "" {
		// Deserialize cachedBook and return
		books := []*models.Category{}
		err := json.Unmarshal([]byte(cachedBook), &books)
		if err != nil {
			return nil, err
		}
		return books, nil
	}

	// Calculate the correct offset
	offset := (page - 1) * limit

	var books []*models.Category
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
