package entities

import "time"

type Order struct {
	OrderID     string    `gorm:"primaryKey;autoIncrement;"`
	MenuID      string    `gorm:"type:varchar(64);not null;"`
	CustomerID  string    `gorm:"type:varchar(64);not null;"`
	OrderAmount uint      `gorm:"not null;"`
	OrderStatus string    `gorm:"type:varchar(64);not null;"`
	CreateAt    time.Time `gorm:"not null;autoCreateTime;"`
}
