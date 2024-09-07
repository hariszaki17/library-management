package repository

import (
	"context"

	"github.com/hariszaki17/library-management/proto/cache"
	"github.com/hariszaki17/library-management/user-service/models"
	"gorm.io/gorm"
)

type BorrowingRecordRepository interface {
	CreateBorrowingRecordWithCtx(tx *gorm.DB, borrowingRecord *models.BorrowingRecord) (*models.BorrowingRecord, error)
	UpdateBorrowingRecordWithCtx(tx *gorm.DB, existingModel *models.BorrowingRecord, updatedFields map[string]interface{}) (*models.BorrowingRecord, error)
	GetBorrowingRecordByIDWithCtx(tx *gorm.DB, id uint) (*models.BorrowingRecord, error)
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

func (r *borrowingRecordRepository) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (r *borrowingRecordRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (r *borrowingRecordRepository) Begin(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Begin()
}
