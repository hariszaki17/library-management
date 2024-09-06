// usecase/user_usecase.go
package usecase

import (
	"context"

	"github.com/hariszaki17/library-management/book-service/models"
	"github.com/hariszaki17/library-management/book-service/repository"
)

type BookUsecase interface {
	GetBooks(ctx context.Context, page, limit int) ([]*models.Book, error)
}

type bookUsecase struct {
	bookRepo repository.BookRepository
}

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	return &bookUsecase{bookRepo: bookRepo}
}

func (u *bookUsecase) GetBooks(ctx context.Context, page, limit int) ([]*models.Book, error) {
	return u.bookRepo.GetBooks(ctx, page, limit)
}
