package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Phone     string    `gorm:"size:20;uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Addresses []Address `gorm:"foreignKey:UserID"`
}

type Address struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"index"`
	Street  string `gorm:"size:255;not null"`
	City    string `gorm:"size:100;not null"`
	State   string `gorm:"size:100;not null"`
	Country string `gorm:"size:100;not null"`
	ZipCode string `gorm:"size:20;not null"`
}
