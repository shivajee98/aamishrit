package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"unique;not null" json:"name"`
	Description string    `json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
