package usecase

import (
	"context"
	"errors"

	"github.com/hariszaki17/library-management/book-service/models"
	"github.com/hariszaki17/library-management/book-service/repository"
	pbUser "github.com/hariszaki17/library-management/proto/gen/user/proto"
	"gorm.io/gorm"
)

type BookUsecase interface {
	GetBooks(ctx context.Context, page, limit int, query string) ([]*models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	UpdateBook(ctx context.Context, id uint, updateValues map[string]any) (*models.Book, error)
	DeleteBook(ctx context.Context, id uint) error
	BorrowBookByID(ctx context.Context, id uint) error
	ReturnBookByID(ctx context.Context, id uint) error
	GetBookRecommendation(ctx context.Context) ([]*models.BookRecommendation, error)
}

type bookUsecase struct {
	bookRepo repository.BookRepository
	userRPC  pbUser.UserServiceClient
}

func NewBookUsecase(bookRepo repository.BookRepository, userRPC *pbUser.UserServiceClient) BookUsecase {
	return &bookUsecase{bookRepo: bookRepo, userRPC: *userRPC}
}

func (u *bookUsecase) GetBooks(ctx context.Context, page, limit int, query string) ([]*models.Book, error) {
	return u.bookRepo.GetBooks(ctx, page, limit, query)
}

func (u *bookUsecase) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	tx := u.bookRepo.Begin(ctx)
	res, err := u.bookRepo.CreateBookWithCtx(tx, book)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return res, nil
}

func (u *bookUsecase) UpdateBook(ctx context.Context, id uint, updatedValues map[string]any) (*models.Book, error) {
	tx := u.bookRepo.Begin(ctx)

	res, err := u.bookRepo.UpdateBookWithCtx(tx, &models.Book{
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

func (u *bookUsecase) DeleteBook(ctx context.Context, id uint) error {
	tx := u.bookRepo.Begin(ctx)
	err := u.bookRepo.DeleteBookWithCtx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *bookUsecase) BorrowBookByID(ctx context.Context, id uint) error {
	tx := u.bookRepo.Begin(ctx)

	book, err := u.bookRepo.GetBookByIDWithCtx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if book.Stock < 1 {
		tx.Rollback()
		return errors.New("book out of stock")
	}

	updatedStock := book.Stock - 1
	_, err = u.bookRepo.UpdateBookWithCtx(tx, book, map[string]interface{}{
		"stock": updatedStock,
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *bookUsecase) ReturnBookByID(ctx context.Context, id uint) error {
	tx := u.bookRepo.Begin(ctx)

	book, err := u.bookRepo.GetBookByIDWithCtx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	updatedStock := book.Stock + 1
	_, err = u.bookRepo.UpdateBookWithCtx(tx, book, map[string]interface{}{
		"stock": updatedStock,
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *bookUsecase) GetBookRecommendation(ctx context.Context) ([]*models.BookRecommendation, error) {
	var bookRecommendation []*models.BookRecommendation
	borrowingCount, err := u.userRPC.GetBorrowingCount(ctx, &pbUser.GetBorrowingCountRequest{})
	if err != nil {
		return nil, err
	}

	if len(borrowingCount.BorrowingCount) < 1 {
		return bookRecommendation, nil
	}

	var ids []uint
	for _, bc := range borrowingCount.BorrowingCount {
		ids = append(ids, uint(bc.BookId))
		bookRecommendation = append(bookRecommendation, &models.BookRecommendation{
			Book: &models.Book{
				Model: gorm.Model{
					ID: uint(bc.BookId),
				},
			},
			BorrowedCount: uint(bc.Count),
		})
	}

	books, err := u.bookRepo.GetBookByIds(ctx, ids)
	if err != nil {
		return bookRecommendation, nil
	}

out:
	for _, br := range bookRecommendation {
		for _, book := range books {
			if br.Book.ID == book.ID {
				br.Book = book
			}
			continue out
		}
	}

	return bookRecommendation, nil
}
