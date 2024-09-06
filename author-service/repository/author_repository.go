package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hariszaki17/library-management/author-service/models"
	"github.com/hariszaki17/library-management/proto/cache"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	GetAuthors(ctx context.Context, page, limit int) ([]*models.Author, error)
	CreateAuthorWithCtx(tx *gorm.DB, author *models.Author) (*models.Author, error)
	UpdateAuthorWithCtx(tx *gorm.DB, existingModel *models.Author, updatedFields map[string]interface{}) (*models.Author, error)
	DeleteAuthorWithCtx(tx *gorm.DB, id uint) error
	Begin(ctx context.Context) *gorm.DB
	Commit(tx *gorm.DB) *gorm.DB
	Rollback(tx *gorm.DB) *gorm.DB
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
		authors := []*models.Author{}
		err := json.Unmarshal([]byte(cachedAuthor), &authors)
		if err != nil {
			return nil, err
		}
		return authors, nil
	}

	// Calculate the correct offset
	offset := (page - 1) * limit

	var authors []*models.Author
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&authors).Error; err != nil {
		return nil, err
	}

	// Cache the result
	serializedAuthors, err := json.Marshal(authors)
	if err != nil {
		return nil, err
	}
	err = r.cache.Set(cacheKey, string(serializedAuthors), time.Minute)
	if err != nil {
		log.Printf("Failed to set cache: %v", err)
	}
	return authors, nil
}

func (r *authorRepository) CreateAuthorWithCtx(tx *gorm.DB, author *models.Author) (*models.Author, error) {
	if err := tx.Create(&author).Error; err != nil {
		return nil, err
	}

	return author, nil
}

func (r *authorRepository) UpdateAuthorWithCtx(tx *gorm.DB, existingModel *models.Author, updatedFields map[string]interface{}) (*models.Author, error) {
	if err := tx.Model(&existingModel).Updates(updatedFields).Error; err != nil {
		return nil, err
	}

	return existingModel, nil
}

func (r *authorRepository) DeleteAuthorWithCtx(tx *gorm.DB, id uint) error {
	if err := tx.Delete(&models.Author{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *authorRepository) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (r *authorRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (r *authorRepository) Begin(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}
