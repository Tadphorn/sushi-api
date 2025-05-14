package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	_menuModel "sushi-api/pkg/menu/model"
	_menuService "sushi-api/pkg/menu/service"
	"time"

	"github.com/labstack/echo/v4"
)

type MenuControllerImpl struct {
	menuService _menuService.MenuService
}

func NewMenuControllerImpl(menuService _menuService.MenuService) MenuController {
	return &MenuControllerImpl{menuService}
}

func (c *MenuControllerImpl) CreateMenu(pctx echo.Context) error {
	menu := new(_menuModel.MenuReq)

	menu.MenuName = pctx.FormValue("menu_name")

	price, err := strconv.Atoi(pctx.FormValue("menu_price"))
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid price format",
		})
	}
	menu.MenuPrice = uint(price)
	menu.CategoryID = pctx.FormValue("category_id")

	// Upload file
	uniqueFilename, err := c.uploadFile(pctx)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to upload file",
		})
	}
	menu.MenuImage = uniqueFilename

	// Call Service
	menuRes, err := c.menuService.CreateMenu(menu)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create menu",
		})
	}

	return pctx.JSON(http.StatusCreated, menuRes)
}

//-------------------------------------------------------------------------------------------------

func (c *MenuControllerImpl) uploadFile(pctx echo.Context) (string, error) {
	file, err := pctx.FormFile("menu_image")
	if err != nil {
		return "", err
	}
	// test opening the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Customize file name
	timestamp := time.Now().UnixNano() // nano = very precise
	ext := filepath.Ext(file.Filename) // keeps original extension
	uniqueFilename := fmt.Sprintf("menu_%d%s", timestamp, ext)

	// Destination
	dstPath := "uploads/" + uniqueFilename
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return uniqueFilename, nil
}
