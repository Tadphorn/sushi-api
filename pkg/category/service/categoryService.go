package service

import (
	_categoryModel "sushi-api/pkg/category/model"
)

type CategoryService interface {
	CreateCategory(category *_categoryModel.CategoryReq) (*_categoryModel.Category, error)
}
