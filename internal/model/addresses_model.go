package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID            uint   `gorm:"index"`
	Street            string `gorm:"size:255;not null"`
	City              string `gorm:"size:100;not null"`
	State             string `gorm:"size:100;not null"`
	Country           string `gorm:"size:100;not null"`
	ZipCode           string `gorm:"size:20;not null"`
	IsDefaultShipping bool   `gorm:"default:false"`
	IsDefaultBilling  bool   `gorm:"default:false"`
}
