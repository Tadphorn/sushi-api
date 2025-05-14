package repository

import "sushi-api/entities"

type MenuRepository interface {
	CreateMenu(menu *entities.Menu) (*entities.Menu, error)
}
