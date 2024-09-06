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
	CreateCategoryWithCtx(tx *gorm.DB, category *models.Category) (*models.Category, error)
	UpdateCategoryWithCtx(tx *gorm.DB, existingModel *models.Category, updatedFields map[string]interface{}) (*models.Category, error)
	DeleteCategoryWithCtx(tx *gorm.DB, id uint) error
	Begin(ctx context.Context) *gorm.DB
	Commit(tx *gorm.DB) *gorm.DB
	Rollback(tx *gorm.DB) *gorm.DB
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
	cachedCategory, err := r.cache.Get(cacheKey)
	if err == nil && cachedCategory != "" {
		// Deserialize cachedCategory and return
		categories := []*models.Category{}
		err := json.Unmarshal([]byte(cachedCategory), &categories)
		if err != nil {
			return nil, err
		}
		return categories, nil
	}

	// Calculate the correct offset
	offset := (page - 1) * limit

	var categories []*models.Category
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}

	// Cache the result
	serializedCategorys, err := json.Marshal(categories)
	if err != nil {
		return nil, err
	}
	err = r.cache.Set(cacheKey, string(serializedCategorys), time.Minute)
	if err != nil {
		log.Printf("Failed to set cache: %v", err)
	}
	return categories, nil
}

func (r *categoryRepository) CreateCategoryWithCtx(tx *gorm.DB, category *models.Category) (*models.Category, error) {
	if err := tx.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategoryWithCtx(tx *gorm.DB, existingModel *models.Category, updatedFields map[string]interface{}) (*models.Category, error) {
	if err := tx.Model(&existingModel).Updates(updatedFields).Error; err != nil {
		return nil, err
	}

	return existingModel, nil
}

func (r *categoryRepository) DeleteCategoryWithCtx(tx *gorm.DB, id uint) error {
	if err := tx.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (r *categoryRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (r *categoryRepository) Begin(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}
