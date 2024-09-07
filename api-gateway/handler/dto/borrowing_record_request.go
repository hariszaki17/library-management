package dto

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
