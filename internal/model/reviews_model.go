package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID  uint   `gorm:"index"`
	CustomerID uint   `gorm:"index"`
	Rating     uint   `json:"rating"`
	Comment    string `json:"comment"`
}
