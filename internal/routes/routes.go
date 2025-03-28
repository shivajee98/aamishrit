package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	public := app.Group("/public")
	protected := app.Group("/protected")

	public.Get("/")
	protected.Get("/pro")
}
