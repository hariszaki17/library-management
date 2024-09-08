package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hariszaki17/library-management/proto/cache"
	"github.com/hariszaki17/library-management/user-service/models"
	"gorm.io/gorm"
)

type BorrowingRecordRepository interface {
	CreateBorrowingRecordWithCtx(tx *gorm.DB, borrowingRecord *models.BorrowingRecord) (*models.BorrowingRecord, error)
	UpdateBorrowingRecordWithCtx(tx *gorm.DB, existingModel *models.BorrowingRecord, updatedFields map[string]interface{}) (*models.BorrowingRecord, error)
	GetBorrowingRecordByIDWithCtx(tx *gorm.DB, id uint) (*models.BorrowingRecord, error)
	GetBorrowingCount(ctx context.Context) ([]*models.BorrowingCount, error)
	GetBorrowingRecords(ctx context.Context, page, limit int) ([]*models.BorrowingRecord, error)
	Begin(ctx context.Context) *gorm.DB
	Commit(tx *gorm.DB) *gorm.DB
	Rollback(tx *gorm.DB) *gorm.DB
}

type borrowingRecordRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewBorrowingRecordRepository(db *gorm.DB, cache *cache.Cache) BorrowingRecordRepository {
	return &borrowingRecordRepository{db: db, cache: cache}
}

func (r *borrowingRecordRepository) CreateBorrowingRecordWithCtx(tx *gorm.DB, borrowingRecord *models.BorrowingRecord) (*models.BorrowingRecord, error) {
	if err := tx.Create(&borrowingRecord).Error; err != nil {
		return nil, err
	}

	return borrowingRecord, nil
}

func (r *borrowingRecordRepository) UpdateBorrowingRecordWithCtx(tx *gorm.DB, existingModel *models.BorrowingRecord, updatedFields map[string]interface{}) (*models.BorrowingRecord, error) {
	if err := tx.Model(&existingModel).Updates(updatedFields).Error; err != nil {
		return nil, err
	}

	return existingModel, nil
}

func (r *borrowingRecordRepository) GetBorrowingRecordByIDWithCtx(tx *gorm.DB, id uint) (*models.BorrowingRecord, error) {
	var borrowingRecord *models.BorrowingRecord
	if err := tx.First(&borrowingRecord, id).Error; err != nil {
		return nil, err
	}

	return borrowingRecord, nil
}

func (r *borrowingRecordRepository) GetBorrowingCount(ctx context.Context) ([]*models.BorrowingCount, error) {
	var result []*models.BorrowingCount

	err := r.db.WithContext(ctx).Model(&models.BorrowingRecord{}).
		Select("book_id, count(book_id) as count").
		Group("book_id").
		Order("count(book_id) desc").
		Limit(3).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *borrowingRecordRepository) GetBorrowingRecords(ctx context.Context, page, limit int) ([]*models.BorrowingRecord, error) {
	cacheKey := fmt.Sprintf("GetBorrowingRecords,page:%d,limit:%d", page, limit)
	cachedBorrowingRecord, err := r.cache.Get(cacheKey)
	if err == nil && cachedBorrowingRecord != "" {
		// Deserialize cachedBorrowingRecord and return
		borrowingRecords := []*models.BorrowingRecord{}
		err := json.Unmarshal([]byte(cachedBorrowingRecord), &borrowingRecords)
		if err != nil {
			return nil, err
		}
		return borrowingRecords, nil
	}

	// Calculate the correct offset
	offset := (page - 1) * limit

	var borrowingRecords []*models.BorrowingRecord
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&borrowingRecords).Error; err != nil {
		return nil, err
	}

	// Cache the result
	serializedBorrowingRecords, err := json.Marshal(borrowingRecords)
	if err != nil {
		return nil, err
	}
	err = r.cache.Set(cacheKey, string(serializedBorrowingRecords), time.Minute)
	if err != nil {
		log.Printf("Failed to set cache: %v", err)
	}
	return borrowingRecords, nil
}

func (r *borrowingRecordRepository) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (r *borrowingRecordRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (r *borrowingRecordRepository) Begin(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}
