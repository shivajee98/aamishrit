package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/handlers"
)

func Setup(app *fiber.App, productHandler *handlers.ProductHandler, cartHandler *handlers.CartHandler) {
	api := app.Group("/api")

	// Product Routes
	product := api.Group("/product")
	product.Get("/:id", productHandler.GetProductByID)
	product.Get("/", productHandler.ListProducts)

	// Cart Routes
	cart := api.Group("/cart")
	cart.Post("/", cartHandler.AddToCart)
	cart.Get("/:user_id", cartHandler.GetCart)
	cart.Delete("/:cart_id", cartHandler.RemoveFromCart)
	cart.Delete("/clear/:user_id", cartHandler.ClearCart)

}
