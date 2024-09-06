package dto

import (
	"fmt"

	pbCategory "github.com/hariszaki17/library-management/proto/gen/category/proto"
	"github.com/hariszaki17/library-management/proto/utils"
	"google.golang.org/protobuf/types/known/structpb"
)

// GetCategoriesResponse represents the response structure for getting categorys
// @Description A list of categories
// @Example {"categories": [{"id": 1, "name": "Category Name"}]}
type GetCategoriesResponse struct {
	Categories []*Category `json:"categories"`
}

// Category represents the category information
// @Description Category details
// @Example {"id": 1, "name": "Category Name"}
type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CreateCategoryResponse represents the response structure for creating a category
// @Description A success message for category creation
// @Example {"message": "Category created successfully"}
type CreateCategoryResponse struct {
	Message string `json:"message"`
}

// UpdateCategoryResponse represents the response structure for updating a category
// @Description A success message for category update
// @Example {"message": "Category updated successfully"}
type UpdateCategoryResponse struct {
	Message string `json:"message"`
}

// DeleteCategoryResponse represents the response structure for deleting a category
// @Description A success message for category deletion
// @Example {"message": "Category deleted successfully"}
type DeleteCategoryResponse struct {
	Message string `json:"message"`
}

func ToGetCategoriesResponse(categories []*pbCategory.Category) GetCategoriesResponse {
	var res []*Category
	for _, category := range categories {
		res = append(res, &Category{
			ID:   uint(category.Id),
			Name: category.Name,
		})
	}
	return GetCategoriesResponse{
		Categories: res,
	}
}

func ToCreateCategoryResponse(message string) CreateCategoryResponse {
	return CreateCategoryResponse{
		Message: message,
	}
}

func ToUpdateCategoryResponse(message string) UpdateCategoryResponse {
	return UpdateCategoryResponse{
		Message: message,
	}
}

func ToDeleteCategoryResponse(message string) DeleteCategoryResponse {
	return DeleteCategoryResponse{
		Message: message,
	}
}

// CreateCategoryRequest represents the request structure for creating a cateogry
// @Description A request structure for creating a cateogry
// @Example {"name": "Category Name"}
type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required" example:"Category Name"`
}

// UpdateCategoryRequest represents the request structure for updating a cateogry
// @Description A request structure for updating a cateogry
// @Example {"name": "Category Name Updated"}
type UpdateCategoryRequest struct {
	Name *string `json:"name" validate:"required" example:"Category Name Updated"`
}

func CreateCategoryRPCRequest(req *CreateCategoryRequest) *pbCategory.CreateCategoryRequest {
	return &pbCategory.CreateCategoryRequest{
		Name: req.Name,
	}
}

func UpdateCategoryRPCRequest(id uint, req UpdateCategoryRequest) (*pbCategory.UpdateCategoryRequest, error) {
	reqMap, err := utils.StructToMap(req)
	if err != nil {
		return nil, err
	}

	structData, err := structpb.NewStruct(reqMap)
	if err != nil {
		return nil, fmt.Errorf("failed to create structpb.Struct: %v", err)
	}

	return &pbCategory.UpdateCategoryRequest{
		Id:   uint64(id),
		Data: structData,
	}, nil
}

func DeleteCategoryRPCRequest(id uint) *pbCategory.DeleteCategoryRequest {
	return &pbCategory.DeleteCategoryRequest{
		Id: uint64(id),
	}
}
