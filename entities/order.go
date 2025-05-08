package entities

import "time"

type Order struct {
	OrderID     string    `gorm:"type:uuid;primary_key;"`
	MenuID      string    `gorm:"type:varchar(255);not null;"`
	CustomerID  string    `gorm:"type:varchar(255);not null;"`
	OrderAmount uint      `gorm:"not null;"`
	OrderStatus string    `gorm:"type:varchar(64);not null;"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
}
