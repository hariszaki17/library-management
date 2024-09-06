package dto

import (
	pbBook "github.com/hariszaki17/library-management/proto/gen/book/proto"
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
