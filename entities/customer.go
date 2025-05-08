package entities

import "time"

type Customer struct {
	CustomerID    string    `gorm:"type:uuid;primary_key;"`
	CustomerTable string    `gorm:"type:varchar(64);not null;"`
	CreatedAt     time.Time `gorm:"not null;autoCreateTime;"`
}
