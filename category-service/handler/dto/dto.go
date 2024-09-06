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

func ToCreateCategoryResponse(message string) *pb.CreateCategoryResponse {
	return &pb.CreateCategoryResponse{
		Message: message,
	}
}

func ToUpdateCategoryResponse(message string) *pb.UpdateCategoryResponse {
	return &pb.UpdateCategoryResponse{
		Message: message,
	}
}

func ToDeleteCategoryResponse(message string) *pb.DeleteCategoryResponse {
	return &pb.DeleteCategoryResponse{
		Message: message,
	}
}
