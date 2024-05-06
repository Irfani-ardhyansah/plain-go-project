package helper

import (
	"golang-api-udemy/model/domain"
	"golang-api-udemy/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryReponses []web.CategoryResponse
	for _, category := range categories {
		categoryReponses = append(categoryReponses, ToCategoryResponse(category))
	}

	return categoryReponses
}
