package entities

import (
	_categoryModel "sushi-api/pkg/category/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	CategoryID   string    `gorm:"type:uuid;primary_key;"`
	CategoryName string    `gorm:"type:varchar(64);not null;"`
	CategoryNo   int       `gorm:"type:int;"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.CategoryID = uuid.New().String()
	return
}

func (c *Category) ToModelCategory() *_categoryModel.Category {
	return &_categoryModel.Category{
		CategoryID:   c.CategoryID,
		CategoryName: c.CategoryName,
		CategoryNo:   c.CategoryNo,
	}
}
