package entities

import (
	"time"

	_menuModel "sushi-api/pkg/menu/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	MenuID     string    `gorm:"type:uuid;primary_key;"`
	MenuName   string    `gorm:"type:varchar(64);not null;"`
	MenuPrice  uint      `gorm:"not null;"`
	MenuImage  string    `gorm:"type:varchar(255);default:null;"`
	CategoryID string    `gorm:"type:uuid;not null;"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;"`
}

func (m *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	m.MenuID = uuid.New().String()
	return
}

func (m *Menu) ToModelMenu() *_menuModel.Menu {
	return &_menuModel.Menu{
		MenuID:     m.MenuID,
		MenuName:   m.MenuName,
		MenuPrice:  m.MenuPrice,
		MenuImage:  m.MenuImage,
		CategoryID: m.CategoryID,
	}
}
