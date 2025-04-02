package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/handlers"
)

func Setup(app *fiber.App, productHandler *handlers.ProductHandler) {
	api := app.Group("/api")

	

	// Product Routes
	product := api.Group("/product")
	product.Get("/:id", productHandler.GetProductByID)
	product.Get("/", productHandler.ListProducts)
}
