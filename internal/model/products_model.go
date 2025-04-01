package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"image"`
	Stock       int      `gorm:"not null" json:"stock"`
	Reviews     []Review `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
}
