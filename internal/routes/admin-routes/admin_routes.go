package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/middleware"
)

func SetupAdminRoutes(app *fiber.App, deps AdminDeps) {
	cfg := config.LoadEnv()
	clerkKey := cfg.ClerkSecretKey

	admin := app.Group("/admin", middleware.ClerkMiddleware(clerkKey), middleware.ClerkAdminMiddleware(clerkKey))

	products := admin.Group("/products")
	// products.Get("/", deps.ProductHandler.GetAllProducts) // ← add this line
	products.Post("/", deps.ProductHandler.CreateProduct)
	products.Put("/:id", deps.ProductHandler.UpdateProduct)
	products.Delete("/:id", deps.ProductHandler.DeleteProduct)

	category := admin.Group("/category")
	category.Post("/", deps.CategoryHandler.CreateCategory)
	category.Get("/:id", deps.CategoryHandler.GetCategoryByID)
	category.Put("/:id", deps.CategoryHandler.UpdateCategory)
	category.Delete("/:id", deps.CategoryHandler.DeleteCategory)
	// Future: Add order mgmt, user banning, refund API, etc.
}
