package dto

import (
	"github.com/hariszaki17/library-management/author-service/models"
	pb "github.com/hariszaki17/library-management/proto/gen/author/proto"
)

func ToGetAuthorsResponse(authors []*models.Author) *pb.GetAuthorsResponse {
	var res []*pb.Author

	for _, author := range authors {
		res = append(res, &pb.Author{
			Id:        uint64(author.ID),
			FirstName: author.FirstName,
			LastName:  author.LastName,
			Biography: author.Biography,
			Birthdate: author.BirthDate.String(),
		})
	}

	return &pb.GetAuthorsResponse{
		Authors: res,
	}
}

func ToCreateAuthorResponse(message string) *pb.CreateAuthorResponse {
	return &pb.CreateAuthorResponse{
		Message: message,
	}
}

func ToUpdateAuthorResponse(message string) *pb.UpdateAuthorResponse {
	return &pb.UpdateAuthorResponse{
		Message: message,
	}
}

func ToDeleteAuthorResponse(message string) *pb.DeleteAuthorResponse {
	return &pb.DeleteAuthorResponse{
		Message: message,
	}
}
