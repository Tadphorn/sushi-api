package controller

import (
	"github.com/labstack/echo/v4"
)

type MenuController interface {
	CreateMenu(pctx echo.Context) error
}
