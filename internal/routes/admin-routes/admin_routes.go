package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/middleware"
)

func SetupAdminRoutes(app *fiber.App, deps AdminDeps) {
	cfg := config.LoadEnv()
	clerkKey := cfg.ClerkSecretKey

	admin := app.Group("/admin", middleware.ClerkMiddleware(clerkKey), middleware.AdminMiddleware())

	products := admin.Group("/products")
	products.Post("/", deps.ProductHandler.CreateProduct)
	products.Put("/:id", deps.ProductHandler.UpdateProduct)
	products.Delete("/:id", deps.ProductHandler.DeleteProduct)

	// Future: Add order mgmt, user banning, refund API, etc.
}
