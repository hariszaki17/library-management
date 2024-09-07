package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hariszaki17/library-management/book-service/models"
	"github.com/hariszaki17/library-management/proto/cache"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetBooks(ctx context.Context, page, limit int) ([]*models.Book, error)
	CreateBookWithCtx(tx *gorm.DB, book *models.Book) (*models.Book, error)
	UpdateBookWithCtx(tx *gorm.DB, existingModel *models.Book, updatedFields map[string]interface{}) (*models.Book, error)
	DeleteBookWithCtx(tx *gorm.DB, id uint) error
	GetBookByIDWithCtx(tx *gorm.DB, id uint) (*models.Book, error)
	Begin(ctx context.Context) *gorm.DB
	Commit(tx *gorm.DB) *gorm.DB
	Rollback(tx *gorm.DB) *gorm.DB
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

func (r *bookRepository) CreateBookWithCtx(tx *gorm.DB, book *models.Book) (*models.Book, error) {
	if err := tx.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepository) UpdateBookWithCtx(tx *gorm.DB, existingModel *models.Book, updatedFields map[string]interface{}) (*models.Book, error) {
	if err := tx.Model(&existingModel).Updates(updatedFields).Error; err != nil {
		return nil, err
	}

	return existingModel, nil
}

func (r *bookRepository) DeleteBookWithCtx(tx *gorm.DB, id uint) error {
	if err := tx.Delete(&models.Book{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *bookRepository) GetBookByIDWithCtx(tx *gorm.DB, id uint) (*models.Book, error) {
	var book *models.Book
	if err := tx.First(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepository) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (r *bookRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (r *bookRepository) Begin(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}
