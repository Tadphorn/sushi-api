package repository

import (
	"sushi-api/databases"
	"sushi-api/entities"

	"github.com/labstack/echo/v4"
)

type MenuRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewMenuRepositoryImpl(db databases.Database, logger echo.Logger) MenuRepository {
	return &MenuRepositoryImpl{db, logger}
}

func (r *MenuRepositoryImpl) CreateMenu(menu *entities.Menu) (*entities.Menu, error) {
	result := r.db.Connect().Create(menu)
	if result.Error != nil {
		r.logger.Error(result.Error)
		return nil, result.Error
	}

	return menu, nil
}
