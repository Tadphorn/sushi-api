package service

import (
	_categoryModel "sushi-api/pkg/category/model"
)

type CategoryService interface {
	CreateCategory(category *_categoryModel.CategoryReq) (*_categoryModel.Category, error)
	EditCategory(categoryID string, newCategory *_categoryModel.CategoryReq) error
	GetAllCategory() ([]*_categoryModel.Category, error)
	DeleteCategory(categoryID string) error
}
