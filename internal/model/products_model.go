package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Images      []string `json:"images" gorm:"type:text[]"`
	Stock       int      `gorm:"not null" json:"stock"`
	Categories  string   `gorm:"not null" json:"category"`
	Reviews     []Review `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
