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
