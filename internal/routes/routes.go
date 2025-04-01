package routes

import (
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/middleware"
)

func Setup(app *fiber.App, jwkClient *jwks.Client, jwkStore *middleware.JWKStore, handler *handlers.Handler) {
	// public routes
	public := app.Group("/public")
	public.Get("public")

	// protected routes
	protected := app.Group("/protected")
	protected.Get("/pro")
}
