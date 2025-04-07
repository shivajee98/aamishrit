package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name                   string     `gorm:"size:255;not null"`
	Phone                  string     `gorm:"size:20;uniqueIndex;not null"`
	Password               string     `gorm:"not null"`
	Addresses              []Address  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Carts                  []Cart     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Reviews                []Review   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Orders                 []Order    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Payments               []Payment  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	DefaultShippingAddress Address    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL;"`
	DefaultBillingAddress  Address    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL;"`
	Wishlists              []Wishlist `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
