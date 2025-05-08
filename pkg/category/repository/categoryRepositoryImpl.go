package repository

import (
	"sushi-api/databases"
	"sushi-api/entities"
	_categoryModel "sushi-api/pkg/category/model"

	"github.com/labstack/echo/v4"
)

type CategoryRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewCategoryRepositoryImpl(db databases.Database, logger echo.Logger) CategoryRepository {
	return &CategoryRepositoryImpl{db, logger}
}

func (r *CategoryRepositoryImpl) GetAllCategory() ([]*entities.Category, error) {
	categories := make([]*entities.Category, 0)
	query := r.db.Connect().Model(&entities.Category{}).Order("category_no").Find(&categories)

	if query.Error != nil {
		r.logger.Error(query.Error)
		return nil, query.Error
	}

	return categories, nil
}

func (r *CategoryRepositoryImpl) CreateCategory(category *entities.Category) (*entities.Category, error) {
	result := r.db.Connect().Create(category)
	if result.Error != nil {
		r.logger.Error(result.Error)
		return nil, result.Error
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) EditCategory(categoryID string, newCategory *_categoryModel.CategoryReq) error {
	panic("implement me")
}

func (r *CategoryRepositoryImpl) DeleteCategory(categoryID string) error {
	panic("implement me")
}
