package entities

import (
	"time"
)

type Menu struct {
	MenuID    string    `gorm:"type:uuid;primary_key;"`
	MenuName  string    `gorm:"type:varchar(64);not null;"`
	MenuPrice uint      `gorm:"not null;"`
	MenuImage string    `gorm:"type:varchar(255);default:null;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
