package entities

import (
	"time"
)

type Menu struct {
	MenuID    string    `gorm:"primaryKey;autoIncrement;"`
	MenuName  string    `gorm:"type:varchar(64);not null;"`
	MenuPrice uint      `gorm:"not null;"`
	MenuImage string    `gorm:"type:varchar(255);default:null;"`
	CreateAt  time.Time `gorm:"not null;autoCreateTime;"`
	UpdateAt  time.Time `gorm:"not null;autoUpdateTime;"`
}
