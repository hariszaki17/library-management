package dto

import (
	"fmt"

	pbBook "github.com/hariszaki17/library-management/proto/gen/book/proto"
	"github.com/hariszaki17/library-management/proto/utils"
	"google.golang.org/protobuf/types/known/structpb"
)

// GetBooksResponse represents the response structure for getting books
// @Description A list of books
// @Example {"books": [{"id": 1, "title": "Book Title", "author_id": 2, "category_id": 3, "isbn": "1234567890", "published_at": "2024-01-01T00:00:00Z", "stock": 10}]}
type GetBooksResponse struct {
	Books []*Book `json:"books"`
}

// Book represents the book information
// @Description Book details
// @Example {"id": 1, "title": "Book Title", "author_id": 2, "category_id": 3, "isbn": "1234567890", "published_at": "2024-01-01T00:00:00Z", "stock": 10}
type Book struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	AuthorID    uint   `json:"author_id"`
	CategoryID  uint   `json:"category_id"`
	ISBN        string `json:"isbn"`
	PublishedAt string `json:"published_at"`
	Stock       uint   `json:"stock"`
}

// GetBookRecommendationResponse represents the response structure for getting book recommendation
// @Description A list of book recommendation
// @Example {"books": [{"id": 1, "title": "Book Title", "author_id": 2, "category_id": 3, "isbn": "1234567890", "published_at": "2024-01-01T00:00:00Z", "stock": 10}]}
type GetBookRecommendationResponse struct {
	BookRecommendations []*BookRecommendation `json:"books"`
}

// BookRecommendation represents the book information
// @Description BookRecommendation details
// @Example {"id": 1, "title": "BookRecommendation Title", "author_id": 2, "category_id": 3, "isbn": "1234567890", "published_at": "2024-01-01T00:00:00Z", "stock": 10, "borrowed_count": 5}
type BookRecommendation struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	AuthorID      uint   `json:"author_id"`
	CategoryID    uint   `json:"category_id"`
	ISBN          string `json:"isbn"`
	PublishedAt   string `json:"published_at"`
	Stock         uint   `json:"stock"`
	BorrowedCount uint   `json:"borrowed_count"`
}

// CreateBookResponse represents the response structure for creating a book
// @Description A success message for book creation
// @Example {"message": "Book created successfully"}
type CreateBookResponse struct {
	Message string `json:"message"`
}

// UpdateBookResponse represents the response structure for updating a book
// @Description A success message for book update
// @Example {"message": "Book updated successfully"}
type UpdateBookResponse struct {
	Message string `json:"message"`
}

// DeleteBookResponse represents the response structure for deleting a book
// @Description A success message for book deletion
// @Example {"message": "Book deleted successfully"}
type DeleteBookResponse struct {
	Message string `json:"message"`
}

func ToGetBooksResponse(books []*pbBook.Book) GetBooksResponse {
	var res []*Book
	for _, book := range books {
		res = append(res, &Book{
			ID:          uint(book.Id),
			Title:       book.Title,
			AuthorID:    uint(book.AuthorId),
			CategoryID:  uint(book.CategoryId),
			ISBN:        book.Isbn,
			PublishedAt: book.PublishedAt,
			Stock:       uint(book.Stock),
		})
	}
	return GetBooksResponse{
		Books: res,
	}
}

func ToGetBookRecommendationResponse(books []*pbBook.BookRecommendation) GetBookRecommendationResponse {
	var res []*BookRecommendation
	for _, book := range books {
		res = append(res, &BookRecommendation{
			ID:            uint(book.Id),
			Title:         book.Title,
			AuthorID:      uint(book.AuthorId),
			CategoryID:    uint(book.CategoryId),
			ISBN:          book.Isbn,
			PublishedAt:   book.PublishedAt,
			Stock:         uint(book.Stock),
			BorrowedCount: uint(book.BorrowedCount),
		})
	}
	return GetBookRecommendationResponse{
		BookRecommendations: res,
	}
}

func ToCreateBookResponse(message string) CreateBookResponse {
	return CreateBookResponse{
		Message: message,
	}
}

func ToUpdateBookResponse(message string) UpdateBookResponse {
	return UpdateBookResponse{
		Message: message,
	}
}

func ToDeleteBookResponse(message string) DeleteBookResponse {
	return DeleteBookResponse{
		Message: message,
	}
}

// CreateBookRequest represents the request structure for creating a book
// @Description A request structure for creating a book
// @Example {"title": "Book Title", "author_id": 2, "category_id": 3, "isbn": "1234567890", "published_at": "2024-01-01T00:00:00Z", "stock": 10}
type CreateBookRequest struct {
	Title       string `json:"title" validate:"required" example:"Book Title"`
	AuthorID    uint64 `json:"author_id" validate:"required" example:"2"`
	CategoryID  uint64 `json:"category_id" validate:"required" example:"3"`
	ISBN        string `json:"isbn" validate:"required" example:"1234567890"`
	PublishedAt string `json:"published_at" validate:"required" example:"2024-01-01T00:00:00Z"`
	Stock       uint64 `json:"stock" example:"10"`
}

// UpdateBookRequest represents the request structure for updating a book
// @Description A request structure for updating a book
// @Example {"title": "Book Title", "author_id": 2, "category_id": 3, "isbn": "1234567890", "published_at": "2024-01-01T00:00:00Z", "stock": 10}
type UpdateBookRequest struct {
	Title       *string `json:"title" example:"Book Title"`
	AuthorID    *uint64 `json:"author_id" example:"2"`
	CategoryID  *uint64 `json:"category_id" example:"3"`
	ISBN        *string `json:"isbn" example:"1234567890"`
	PublishedAt *string `json:"published_at" example:"2024-01-01T00:00:00Z"`
	Stock       *uint64 `json:"stock" example:"10"`
}

func CreateBookRPCRequest(req *CreateBookRequest) *pbBook.CreateBookRequest {
	return &pbBook.CreateBookRequest{
		Title:       req.Title,
		AuthorId:    req.AuthorID,
		CategoryId:  req.CategoryID,
		Isbn:        req.ISBN,
		PublishedAt: req.PublishedAt,
		Stock:       req.Stock,
	}
}

func UpdateBookRPCRequest(id uint, req UpdateBookRequest) (*pbBook.UpdateBookRequest, error) {
	reqMap, err := utils.StructToMap(req)
	if err != nil {
		return nil, err
	}

	structData, err := structpb.NewStruct(reqMap)
	if err != nil {
		return nil, fmt.Errorf("failed to create structpb.Struct: %v", err)
	}

	return &pbBook.UpdateBookRequest{
		Id:   uint64(id),
		Data: structData,
	}, nil
}

func DeleteBookRPCRequest(id uint) *pbBook.DeleteBookRequest {
	return &pbBook.DeleteBookRequest{
		Id: uint64(id),
	}
}
