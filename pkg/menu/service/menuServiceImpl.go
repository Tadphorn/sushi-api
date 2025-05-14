package service

import (
	"sushi-api/entities"
	_menuModel "sushi-api/pkg/menu/model"
	_menuRepository "sushi-api/pkg/menu/repository"
)

type MenuServiceImpl struct {
	menuRepository _menuRepository.MenuRepository
}

func NewMenuServiceImpl(menuRepository _menuRepository.MenuRepository) MenuService {
	return &MenuServiceImpl{menuRepository}
}

func (s *MenuServiceImpl) CreateMenu(menu *_menuModel.MenuReq) (*_menuModel.Menu, error) {
	menuEntity := &entities.Menu{
		MenuName:   menu.MenuName,
		MenuPrice:  menu.MenuPrice,
		MenuImage:  menu.MenuImage,
		CategoryID: menu.CategoryID,
	}

	menuRes, err := s.menuRepository.CreateMenu(menuEntity)
	if err != nil {
		return nil, err
	}
	return menuRes.ToModelMenu(), nil
}
