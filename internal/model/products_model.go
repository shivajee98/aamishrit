package model

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Price       float64        `json:"price"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"` // FIXED HERE
	Stock       int            `gorm:"not null" json:"stock"`
	Categories  string         `gorm:"not null" json:"category"`
	Reviews     []Review       `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
