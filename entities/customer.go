package entities

import "time"

type Customer struct {
	CustomerID    string    `gorm:"primaryKey;autoIncrement;"`
	CustomerTable string    `gorm:"type:varchar(64);not null;"`
	CreateAt      time.Time `gorm:"not null;autoCreateTime;"`
}
