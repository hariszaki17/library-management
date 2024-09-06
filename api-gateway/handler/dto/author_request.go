package dto

import (
	pbAuthor "github.com/hariszaki17/library-management/proto/gen/author/proto"
)

// GetAuthorsResponse represents the response structure for getting authors
// @Description A list of authors
// @Example {"authors": [{"id": 1, "first_name": "John", "last_name": "Doe", "biography": "An accomplished author.", "birth_date": "1980-01-01"}]}
type GetAuthorsResponse struct {
	Authors []*Author `json:"authors"`
}

// Author represents the author information
// @Description Author details
// @Example {"id": 1, "first_name": "John", "last_name": "Doe", "biography": "An accomplished author.", "birth_date": "1980-01-01"}
type Author struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
	BirthDate string `json:"birth_date"`
}

// ToGetAuthorsResponse converts a slice of pbAuthor.Author to GetAuthorsResponse
func ToGetAuthorsResponse(authors []*pbAuthor.Author) GetAuthorsResponse {
	var res []*Author
	for _, author := range authors {
		res = append(res, &Author{
			ID:        uint(author.Id),
			FirstName: author.FirstName,
			LastName:  author.LastName,
			Biography: author.Biography,
			BirthDate: author.Birthdate,
		})
	}

	return GetAuthorsResponse{
		Authors: res,
	}
}
