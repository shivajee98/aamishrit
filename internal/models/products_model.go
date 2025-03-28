package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	ImageUrl    string `json:"image"` // Cloudinary URL for the product image
	CompanyId   uint   `gorm:"index"`
}
