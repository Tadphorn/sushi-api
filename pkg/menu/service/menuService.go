package service

import (
	_menuModel "sushi-api/pkg/menu/model"
)

type MenuService interface {
	CreateMenu(menu *_menuModel.MenuReq) (*_menuModel.Menu, error)
}
