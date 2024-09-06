package usecase

import (
	"context"

	"github.com/hariszaki17/library-management/category-service/models"
	"github.com/hariszaki17/library-management/category-service/repository"
)

type CategoryUsecase interface {
	GetCategories(ctx context.Context, page, limit int) ([]*models.Category, error)
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
