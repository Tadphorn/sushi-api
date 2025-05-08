package repository

import (
	"sushi-api/entities"
	_categoryModel "sushi-api/pkg/category/model"
)

type CategoryRepository interface {
	GetAllCategory() ([]*entities.Category, error)
	CreateCategory(category *entities.Category) (*entities.Category, error)
	EditCategory(categoryID string, newCategory *_categoryModel.CategoryReq) error
	DeleteCategory(categoryID string) error
}
