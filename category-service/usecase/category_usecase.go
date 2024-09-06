package usecase

import (
	"context"

	"github.com/hariszaki17/library-management/category-service/models"
	"github.com/hariszaki17/library-management/category-service/repository"
	"gorm.io/gorm"
)

type CategoryUsecase interface {
	GetCategories(ctx context.Context, page, limit int) ([]*models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	UpdateCategory(ctx context.Context, id uint, updateValues map[string]any) (*models.Category, error)
	DeleteCategory(ctx context.Context, id uint) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo: categoryRepo}
}

func (u *categoryUsecase) GetCategories(ctx context.Context, page, limit int) ([]*models.Category, error) {
	return u.categoryRepo.GetCategories(ctx, page, limit)
}

func (u *categoryUsecase) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	tx := u.categoryRepo.Begin(ctx)
	res, err := u.categoryRepo.CreateCategoryWithCtx(tx, category)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return res, nil
}

func (u *categoryUsecase) UpdateCategory(ctx context.Context, id uint, updatedValues map[string]any) (*models.Category, error) {
	tx := u.categoryRepo.Begin(ctx)

	res, err := u.categoryRepo.UpdateCategoryWithCtx(tx, &models.Category{
		Model: gorm.Model{
			ID: id,
		},
	}, updatedValues)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return res, nil
}

func (u *categoryUsecase) DeleteCategory(ctx context.Context, id uint) error {
	tx := u.categoryRepo.Begin(ctx)
	err := u.categoryRepo.DeleteCategoryWithCtx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
