package usecase

import (
	"context"

	"github.com/hariszaki17/library-management/author-service/models"
	"github.com/hariszaki17/library-management/author-service/repository"
	"gorm.io/gorm"
)

type AuthorUsecase interface {
	GetAuthors(ctx context.Context, page, limit int) ([]*models.Author, error)
	CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error)
	UpdateAuthor(ctx context.Context, id uint, updateValues map[string]any) (*models.Author, error)
	DeleteAuthor(ctx context.Context, id uint) error
}

type authorUsecase struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorUsecase(authorRepo repository.AuthorRepository) AuthorUsecase {
	return &authorUsecase{authorRepo: authorRepo}
}

func (u *authorUsecase) GetAuthors(ctx context.Context, page, limit int) ([]*models.Author, error) {
	return u.authorRepo.GetAuthors(ctx, page, limit)
}


func (u *authorUsecase) CreateAuthor(ctx context.Context, author *models.Author) (*models.Author, error) {
	tx := u.authorRepo.Begin(ctx)
	res, err := u.authorRepo.CreateAuthorWithCtx(tx, author)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return res, nil
}

func (u *authorUsecase) UpdateAuthor(ctx context.Context, id uint, updatedValues map[string]any) (*models.Author, error) {
	tx := u.authorRepo.Begin(ctx)

	res, err := u.authorRepo.UpdateAuthorWithCtx(tx, &models.Author{
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

func (u *authorUsecase) DeleteAuthor(ctx context.Context, id uint) error {
	tx := u.authorRepo.Begin(ctx)
	err := u.authorRepo.DeleteAuthorWithCtx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
