package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string         `gorm:"unique;not null" json:"name"`
	Description string         `json:"description"`
	Products    []*Product     `gorm:"many2many:product_categories;"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
