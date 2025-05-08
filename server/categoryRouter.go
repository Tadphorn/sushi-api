package server

import (
	_categoryController "sushi-api/pkg/category/controller"
	_categoryRepository "sushi-api/pkg/category/repository"
	_categoryService "sushi-api/pkg/category/service"
)

func (s *echoServer) initCategoryRouter() {
	router := s.app.Group("/category")

	categoryRepository := _categoryRepository.NewCategoryRepositoryImpl(s.db, s.app.Logger)
	categoryService := _categoryService.NewCategoryServiceImpl(categoryRepository)
	categoryController := _categoryController.NewCategoryControllerImpl(categoryService)

	router.POST("/create", categoryController.CreateCategory)
	router.GET("/getall", categoryController.GetAllCategory)
	router.PUT("/edit/:id", categoryController.EditCategory)
}
