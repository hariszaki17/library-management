package dto

import (
	"fmt"

	pbAuthor "github.com/hariszaki17/library-management/proto/gen/author/proto"
	"github.com/hariszaki17/library-management/proto/utils"
	"google.golang.org/protobuf/types/known/structpb"
)

// GetAuthorsResponse represents the response structure for getting authors
// @Description A list of authors
// @Example {"authors": [{"id": 1, "first_name": "John", "last_name": "Doe", "biography": "An accomplished author.", "birth_date": "1980-01-01"}]}
type GetAuthorsResponse struct {
	Authors []*Author `json:"authors"`
}

// CreateAuthorResponse represents the response structure for creating an author
// @Description A success message for author creation
// @Example {"message": "Author created successfully"}
type CreateAuthorResponse struct {
	Message string `json:"message"`
}

// UpdateAuthorResponse represents the response structure for updating an author
// @Description A success message for author update
// @Example {"message": "Author updated successfully"}
type UpdateAuthorResponse struct {
	Message string `json:"message"`
}

// DeleteAuthorResponse represents the response structure for deleting an author
// @Description A success message for author deletion
// @Example {"message": "Author deleted successfully"}
type DeleteAuthorResponse struct {
	Message string `json:"message"`
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

func ToCreateAuthorResponse(message string) CreateAuthorResponse {
	return CreateAuthorResponse{
		Message: message,
	}
}

func ToUpdateAuthorResponse(message string) UpdateAuthorResponse {
	return UpdateAuthorResponse{
		Message: message,
	}
}

func ToDeleteAuthorResponse(message string) DeleteAuthorResponse {
	return DeleteAuthorResponse{
		Message: message,
	}
}

// CreateAuthorRequest represents the request structure for creating an author
// @Description A request structure for creating an author
// @Example {"first_name": "John", "last_name": "Doe", "biography": "An accomplished author.", "birth_date": "1980-01-01"}
type CreateAuthorRequest struct {
	FirstName string `json:"first_name" validate:"required" example:"John"`
	LastName  string `json:"last_name" validate:"required" example:"Doe"`
	Biography string `json:"biography" validate:"required" example:"An accomplished author."`
	BirthDate string `json:"birth_date" validate:"required" example:"1980-01-01"`
}

// UpdateAuthorRequest represents the request structure for updateing an author
// @Description A request structure for updating an author
// @Example {"first_name": "John", "last_name": "Doe", "biography": "An accomplished author.", "birth_date": "1980-01-01"}
type UpdateAuthorRequest struct {
	FirstName *string `json:"first_name" example:"John Updated"`
	LastName  *string `json:"last_name" example:"Doe Update"`
	Biography *string `json:"biography" example:"An accomplished author."`
	BirthDate *string `json:"birth_date" example:"1990-01-01"`
}

func CreateAuthorRPCRequest(req *CreateAuthorRequest) *pbAuthor.CreateAuthorRequest {
	return &pbAuthor.CreateAuthorRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Biography: req.Biography,
		Birthdate: req.BirthDate,
	}
}

func UpdateAuthorRPCRequest(id uint, req UpdateAuthorRequest) (*pbAuthor.UpdateAuthorRequest, error) {
	reqMap, err := utils.StructToMap(req)
	if err != nil {
		return nil, err
	}

	structData, err := structpb.NewStruct(reqMap)
	if err != nil {
		return nil, fmt.Errorf("failed to create structpb.Struct: %v", err)
	}

	return &pbAuthor.UpdateAuthorRequest{
		Id:   uint64(id),
		Data: structData,
	}, nil
}

func DeleteAuthorRPCRequest(id uint) *pbAuthor.DeleteAuthorRequest {
	return &pbAuthor.DeleteAuthorRequest{
		Id: uint64(id),
	}
}
