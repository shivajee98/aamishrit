package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/handlers"
)

func Setup(app *fiber.App, productHandler *handlers.ProductHandler, cartHandler *handlers.CartHandler, reviewHandler *handlers.ReviewHandler) {
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

	// Review Routes
	reviewRoutes := app.Group("/reviews")

	reviewRoutes.Post("/", reviewHandler.AddReview)
	reviewRoutes.Get("/:product_id", reviewHandler.GetReviews)
	reviewRoutes.Put("/:review_id", reviewHandler.UpdateReview)
	reviewRoutes.Delete("/:review_id", reviewHandler.DeleteReview)

}
