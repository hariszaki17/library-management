package dto

import (
	"github.com/hariszaki17/library-management/category-service/models"
	pb "github.com/hariszaki17/library-management/proto/gen/category/proto"
)

func ToGetCategoriesResponse(categories []*models.Category) *pb.GetCategoriesResponse {
	var res []*pb.Category

	for _, category := range categories {
		res = append(res, &pb.Category{
			Id:   uint64(category.ID),
			Name: category.Name,
		})
	}

	return &pb.GetCategoriesResponse{
		Categories: res,
	}
}
