// usecase/user_usecase.go
package usecase

import (
	"context"
	"time"

	pbBook "github.com/hariszaki17/library-management/proto/gen/book/proto"

	"github.com/hariszaki17/library-management/user-service/models"
	"github.com/hariszaki17/library-management/user-service/repository"
)

type BorrowingRecordUsecase interface {
	BorrowBook(ctx context.Context, userID, bookId uint) error
	ReturnBook(ctx context.Context, id uint) error
	GetBorrowingCount(ctx context.Context) ([]*models.BorrowingCount, error)
	GetBorrowingRecords(ctx context.Context, page, limit int) ([]*models.BorrowingRecord, error)
}

type borrowingRecordUsecase struct {
	borrowingRecordRepo repository.BorrowingRecordRepository
	bookRPC             pbBook.BookServiceClient
}

func NewBorrowingRecordUsecase(borrowingRecordRepo repository.BorrowingRecordRepository, bookRPC *pbBook.BookServiceClient) BorrowingRecordUsecase {
	return &borrowingRecordUsecase{borrowingRecordRepo: borrowingRecordRepo, bookRPC: *bookRPC}
}

func (u *borrowingRecordUsecase) BorrowBook(ctx context.Context, userID, bookID uint) error {
	tx := u.borrowingRecordRepo.Begin(ctx)

	_, err := u.borrowingRecordRepo.CreateBorrowingRecordWithCtx(tx, &models.BorrowingRecord{
		UserID:     userID,
		BookID:     bookID,
		BorrowedAt: time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = u.bookRPC.BorrowBookByID(ctx, &pbBook.BorrowBookByIDRequest{
		Id: uint64(bookID),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *borrowingRecordUsecase) ReturnBook(ctx context.Context, id uint) error {
	tx := u.borrowingRecordRepo.Begin(ctx)

	borrowingRecord, err := u.borrowingRecordRepo.GetBorrowingRecordByIDWithCtx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = u.borrowingRecordRepo.UpdateBorrowingRecordWithCtx(tx, borrowingRecord, map[string]interface{}{
		"returned_at": time.Now(),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = u.bookRPC.ReturnBookByID(ctx, &pbBook.ReturnBookByIDRequest{
		Id: uint64(borrowingRecord.BookID),
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *borrowingRecordUsecase) GetBorrowingCount(ctx context.Context) ([]*models.BorrowingCount, error) {
	return u.borrowingRecordRepo.GetBorrowingCount(ctx)
}

func (u *borrowingRecordUsecase) GetBorrowingRecords(ctx context.Context, page, limit int) ([]*models.BorrowingRecord, error) {
	return u.borrowingRecordRepo.GetBorrowingRecords(ctx, page, limit)
}