package dto

import (
	"github.com/hariszaki17/library-management/book-service/models"
	pb "github.com/hariszaki17/library-management/proto/gen/book/proto"
)

func ToGetBooksResponse(books []*models.Book) *pb.GetBooksResponse {
	var res []*pb.Book

	for _, book := range books {
		res = append(res, &pb.Book{
			Id:          uint64(book.ID),
			Title:       book.Title,
			AuthorId:    uint64(book.AuthorID),
			CategoryId:  uint64(book.CategoryID),
			Isbn:        book.ISBN,
			PublishedAt: book.PublishedAt.String(),
			Stock:       uint64(book.Stock),
		})
	}

	return &pb.GetBooksResponse{
		Books: res,
	}
}

func ToCreateBookResponse(message string) *pb.CreateBookResponse {
	return &pb.CreateBookResponse{
		Message: message,
	}
}

func ToUpdateBookResponse(message string) *pb.UpdateBookResponse {
	return &pb.UpdateBookResponse{
		Message: message,
	}
}

func ToDeleteBookResponse(message string) *pb.DeleteBookResponse {
	return &pb.DeleteBookResponse{
		Message: message,
	}
}

func ToBorrowBookByIDResponse(message string) *pb.BorrowBookByIDResponse {
	return &pb.BorrowBookByIDResponse{
		Message: message,
	}
}

func ToReturnBookByIDResponse(message string) *pb.ReturnBookByIDResponse {
	return &pb.ReturnBookByIDResponse{
		Message: message,
	}
}

func ToGetBookRecommendationResponse(books []*models.BookRecommendation) *pb.GetBookRecommendationResponse {
	var res []*pb.BookRecommendation

	for _, book := range books {
		res = append(res, &pb.BookRecommendation{
			Id:            uint64(book.ID),
			Title:         book.Title,
			AuthorId:      uint64(book.AuthorID),
			CategoryId:    uint64(book.CategoryID),
			Isbn:          book.ISBN,
			PublishedAt:   book.PublishedAt.String(),
			Stock:         uint64(book.Stock),
			BorrowedCount: uint64(book.BorrowedCount),
		})
	}

	return &pb.GetBookRecommendationResponse{
		BookRecommendation: res,
	}
}
