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

func (s *CategoryServiceImpl) EditCategory(categoryID string, newCategory *_categoryModel.CategoryReq) error {
	categoryEntity := &entities.Category{
		CategoryName: newCategory.CategoryName,
		CategoryNo:   newCategory.CategoryNo,
	}

	// Check is category exists
	if err := s.categoryRepository.FindByID(categoryID); err != nil {
		return err
	}

	err := s.categoryRepository.EditCategory(categoryID, categoryEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryServiceImpl) GetAllCategory() ([]*_categoryModel.Category, error) {
	categories, err := s.categoryRepository.GetAllCategory()
	if err != nil {
		return nil, err
	}

	// Convert entities to model
	categoriesModel := make([]*_categoryModel.Category, 0)
	for _, category := range categories {
		categoriesModel = append(categoriesModel, category.ToModelCategory())
	}

	return categoriesModel, nil
}

func (s *CategoryServiceImpl) DeleteCategory(categoryID string) error {
	// Check is category exists
	if err := s.categoryRepository.FindByID(categoryID); err != nil {
		return err
	}

	err := s.categoryRepository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}

	return nil
}
