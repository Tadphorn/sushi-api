package model

type (
	Category struct {
		CategoryID   string `json:"category_id"`
		CategoryName string `json:"category_name"`
		CategoryNo   int    `json:"category_no"`
	}

	CategoryReq struct {
		CategoryName string `json:"category_name" validate:"required,max=64"`
		CategoryNo   int    `json:"category_no"`
	}
)
