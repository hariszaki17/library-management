package usecase

import (
	"context"

	"github.com/hariszaki17/library-management/author-service/models"
	"github.com/hariszaki17/library-management/author-service/repository"
)

type AuthorUsecase interface {
	GetAuthors(ctx context.Context, page, limit int) ([]*models.Author, error)
}

type bookUsecase struct {
	bookRepo repository.AuthorRepository
}

func NewAuthorUsecase(bookRepo repository.AuthorRepository) AuthorUsecase {
	return &bookUsecase{bookRepo: bookRepo}
}

func (u *bookUsecase) GetAuthors(ctx context.Context, page, limit int) ([]*models.Author, error) {
	return u.bookRepo.GetAuthors(ctx, page, limit)
}
