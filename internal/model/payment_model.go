package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID  uint    `gorm:"index"`
	OrderID uint    `gorm:"index"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"` // success, failed, pending
}
