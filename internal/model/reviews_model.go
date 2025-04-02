package model

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    uint      `gorm:"index"`
	ProductID uint      `gorm:"index"`
	Rating    int       `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
