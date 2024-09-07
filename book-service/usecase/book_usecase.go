package usecase

import (
	"context"
	"errors"

	"github.com/hariszaki17/library-management/book-service/models"
	"github.com/hariszaki17/library-management/book-service/repository"
	"gorm.io/gorm"
)

type BookUsecase interface {
	GetBooks(ctx context.Context, page, limit int) ([]*models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	UpdateBook(ctx context.Context, id uint, updateValues map[string]any) (*models.Book, error)
	DeleteBook(ctx context.Context, id uint) error
	BorrowBookByID(ctx context.Context, id uint) error
	ReturnBookByID(ctx context.Context, id uint) error
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
