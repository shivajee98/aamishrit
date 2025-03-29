package handlers

import "gorm.io/gorm"

// handler struct to inject dependencies
type Handler struct {
	DB *gorm.DB
}

func Provide(db *gorm.DB) *Handler{
	return &Handler{DB: db}
}
