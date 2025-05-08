package controller

import "github.com/labstack/echo/v4"

type CategoryController interface {
	CreateCategory(pctx echo.Context) error
	GetAllCategory(pctx echo.Context) error
}
