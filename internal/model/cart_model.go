package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint    `gorm:"index"`
	ProductID uint    `gorm:"index"`
	Quantity  int     `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
}
