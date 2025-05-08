package repository

import (
	"sushi-api/entities"
)

type CategoryRepository interface {
	GetAllCategory() ([]*entities.Category, error)
	CreateCategory(category *entities.Category) (*entities.Category, error)
	EditCategory(categoryID string, newCategory *entities.Category) error
	DeleteCategory(categoryID string) error
	FindByID(categoryID string) error
}
