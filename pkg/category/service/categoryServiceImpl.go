package service

import (
	"sushi-api/entities"
	_categoryModel "sushi-api/pkg/category/model"
	_categoryRepository "sushi-api/pkg/category/repository"
)

type CategoryServiceImpl struct {
	categoryRepository _categoryRepository.CategoryRepository
}

func NewCategoryServiceImpl(categoryRepository _categoryRepository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{categoryRepository}
}

func (s *CategoryServiceImpl) CreateCategory(category *_categoryModel.CategoryReq) (*_categoryModel.Category, error) {
	categoryEntity := &entities.Category{
		CategoryName: category.CategoryName,
		CategoryNo:   category.CategoryNo,
	}

	categoryEntityRes, err := s.categoryRepository.CreateCategory(categoryEntity)
	if err != nil {
		return nil, err
	}
	return categoryEntityRes.ToModelCategory(), nil
}

// func (s CategoryServiceImpl) gennerateCategoryID() string {

// }
