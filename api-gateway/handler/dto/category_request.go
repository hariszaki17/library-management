package dto

import pbCategory "github.com/hariszaki17/library-management/proto/gen/category/proto"

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
