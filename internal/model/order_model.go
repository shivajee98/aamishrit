package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID            uint        `gorm:"index"`
	Total             float64     `json:"total"`
	Status            string      `json:"status"`
	ShippingAddressID uint        `gorm:"index"`
	BillingAddressID  uint        `gorm:"index"`
	ShippingAddress   Address     `gorm:"foreignKey:ShippingAddressID;constraint:OnDelete:SET NULL;"`
	BillingAddress    Address     `gorm:"foreignKey:BillingAddressID;constraint:OnDelete:SET NULL;"`
	Items             []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	Payments          []Payment   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"index"`
	ProductID uint    `gorm:"index"`
	Quantity  int     `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`
}
