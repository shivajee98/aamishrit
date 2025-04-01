package model

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}
