package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID          uint    `gorm:"index" json:"user_id"`
	OrderID         uint    `gorm:"index" json:"order_id"`
	Amount          float64 `gorm:"not null" json:"amount"`
	Currency        string  `gorm:"type:varchar(10);default:'INR'" json:"currency"`
	Status          string  `gorm:"type:varchar(20);default:'pending'" json:"status"` // success, failed, pending
	PaymentMethod   string  `gorm:"type:varchar(50)" json:"payment_method"`           // UPI, Card, Wallet, etc.
	TransactionID   string  `gorm:"uniqueIndex" json:"transaction_id"`                // Razorpay Payment ID
	RazorpayOrderID string  `gorm:"uniqueIndex" json:"razorpay_order_id"`
}
