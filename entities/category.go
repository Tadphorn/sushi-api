package entities

import "time"

type Category struct {
	CategoryID   string    `gorm:"primaryKey;autoIncrement;"`
	CategoryName string    `gorm:"type:varchar(64);not null;"`
	CreateAt     time.Time `gorm:"not null;autoCreateTime;"`
}
