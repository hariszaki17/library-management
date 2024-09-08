package dto

import (
	pbUser "github.com/hariszaki17/library-management/proto/gen/user/proto"
)

// UserBorrowBookResponse represents the response structure for borrow a book
// @Description User borrow book response data
// @Example {"message": "Successfully user borrow a book"}
type UserBorrowBookResponse struct {
	Message string `json:"message"`
}

// UserBorrowBookRequest represents the request structure for borrow a book
// @Description A request structure for borrowing a book
// @Example {"book_id": 1}
type UserBorrowBookRequest struct {
	BookID uint `json:"book_id" validate:"required" example:"1"`
}

func ToUserBorrowBookResponse(message string) UserBorrowBookResponse {
	return UserBorrowBookResponse{
		Message: message,
	}
}

// UserReturnBookResponse represents the response structure for return a book
// @Description User return book response data
// @Example {"message": "Successfully user return a book"}
type UserReturnBookResponse struct {
	Message string `json:"message"`
}

// UserReturnBookRequest represents the request structure for return a book
// @Description A request structure for returning a book
// @Example {"id": 1}
type UserReturnBookRequest struct {
	ID uint `json:"id" validate:"required" example:"1"`
}

func ToUserReturnBookResponse(message string) UserReturnBookResponse {
	return UserReturnBookResponse{
		Message: message,
	}
}

// GetBorrowingRecordsResponse represents the response structure for getting borrowing records
// @Description A list of borrowing records
// @Example {"borrowing_records": [{"id": 1, "user_id": 1, "book_id": 1, "borrowed_at": "2024-01-01T00:00:00Z", "returned_at": "2024-02-01T00:00:00Z"}]}
type GetBorrowingRecordsResponse struct {
	BorrowingRecords []*BorrowingRecords `json:"borrowing_records"`
}

// BorrowingRecords represents the borrowing records information
// @Description BorrowingRecords details
// @Example {"id": 1, "user_id": 1, "book_id": 1, "borrowed_at": "2024-01-01T00:00:00Z", "returned_at": "2024-02-01T00:00:00Z"}
type BorrowingRecords struct {
	ID     uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	BookID     uint   `json:"book_id"`
	BorrowedAt string `json:"borrowed_at"`
	ReturnedAt string `json:"returned_at"`
}

func ToGetBorrowingRecordsResponse(borrowingRecords []*pbUser.BorrowingRecord) GetBorrowingRecordsResponse {
	var res []*BorrowingRecords
	for _, br := range borrowingRecords {
		res = append(res, &BorrowingRecords{
			ID:         uint(br.Id),
			UserID:     uint(br.UserId),
			BookID:     uint(br.BookId),
			BorrowedAt: br.BorrowedAt,
			ReturnedAt: br.ReturnedAt,
		})
	}
	return GetBorrowingRecordsResponse{
		BorrowingRecords: res,
	}
}
