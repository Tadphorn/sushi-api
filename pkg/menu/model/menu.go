package model

type (
	Menu struct {
		MenuID     string `json:"menu_id"`
		MenuName   string `json:"menu_name"`
		MenuPrice  uint   `json:"menu_price"`
		MenuImage  string `json:"menu_image"`
		CategoryID string `json:"category_id"`
	}

	MenuReq struct {
		MenuName   string `json:"menu_name" validate:"required,max=64"`
		MenuPrice  uint   `json:"menu_price" validate:"required"`
		MenuImage  string `json:"menu_image" validate:"required"`
		CategoryID string `json:"category_id" validate:"required"`
	}
)
