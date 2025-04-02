package model

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	Name     string     `json:"name"`
	Products []*Product `gorm:"many2many:product_categories;"`
}
