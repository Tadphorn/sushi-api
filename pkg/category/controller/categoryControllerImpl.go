package controller

import (
	"net/http"
	_categoryModel "sushi-api/pkg/category/model"
	_categoryService "sushi-api/pkg/category/service"

	"github.com/labstack/echo/v4"
)

type CategoryControllerImpl struct {
	categoryService _categoryService.CategoryService
}

func NewCategoryControllerImpl(categoryService _categoryService.CategoryService) CategoryController {
	return &CategoryControllerImpl{categoryService}
}

func (c *CategoryControllerImpl) CreateCategory(pctx echo.Context) error {
	categoryReq := new(_categoryModel.CategoryReq)

	// Bind JSON to struct
	if err := pctx.Bind(&categoryReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	// Call Service
	category, err := c.categoryService.CreateCategory(categoryReq)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create category",
		})
	}

	return pctx.JSON(http.StatusCreated, category)
}

func (c *CategoryControllerImpl) GetAllCategory(pctx echo.Context) error {
	categories, err := c.categoryService.GetAllCategory()
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch categories",
		})
	}

	return pctx.JSON(http.StatusOK, categories)
}
