package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID    uint   `gorm:"index"`
	Street    string `json:"street" gorm:"size:255;not null"`
	City      string `json:"city" gorm:"size:100;not null"`
	State     string `json:"state" gorm:"size:100;not null"`
	Country   string `json:"country" gorm:"size:100;not null"`
	ZipCode   string `json:"zipCode" gorm:"size:20;not null"`
	IsDefault bool   `json:"isDefault" gorm:"default:false"`
}
