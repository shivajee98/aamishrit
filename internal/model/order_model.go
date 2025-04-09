package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID            uint        `gorm:"index" json:"user_id"`
	TotalAmount       float64     `json:"total_amount"`
	Status            string      `gorm:"type:varchar(20);default:'pending'" json:"status"`
	ShippingAddressID uint        `gorm:"index" json:"shipping_address_id"` // <- THIS is required
	ShippingAddress   Address     `gorm:"foreignKey:ShippingAddressID;constraint:OnDelete:SET NULL;" json:"shipping_address"`
	Items             []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"items"`
	Payments          []Payment   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"payments"`
	CreatedAt         int64       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         int64       `gorm:"autoUpdateTime" json:"updated_at"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"index" json:"order_id"`
	ProductID uint    `gorm:"index" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product"`
}
