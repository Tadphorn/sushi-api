package server

import (
	_menuController "sushi-api/pkg/menu/controller"
	_menuRepository "sushi-api/pkg/menu/repository"
	_menuService "sushi-api/pkg/menu/service"
)

func (s *echoServer) initMenuRouter() {
	router := s.app.Group("/menu")

	menuRepository := _menuRepository.NewMenuRepositoryImpl(s.db, s.app.Logger)
	menuService := _menuService.NewMenuServiceImpl(menuRepository)
	menuController := _menuController.NewMenuControllerImpl(menuService)

	router.POST("/create", menuController.CreateMenu)

}
