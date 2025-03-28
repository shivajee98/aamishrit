package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/handlers"
)

func Setup(app *fiber.App, handler *handlers.Handler) {
	public := app.Group("/public")
	protected := app.Group("/protected")

	public.Get("/")
	protected.Get("/pro")
}
