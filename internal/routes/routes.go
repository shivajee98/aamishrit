package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/middleware"
)

func Setup(app *fiber.App, userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler, cartHandler *handlers.CartHandler, reviewHandler *handlers.ReviewHandler) {
	api := app.Group("/api")
	// protected route

	// env loading
	cfg := config.LoadEnv()
	ClerkSecretKey := cfg.ClerkSecretKey

	protected := app.Group("/user", middleware.ClerkMiddleware(ClerkSecretKey))

	// User Routes
	user := protected.Group("/user")
	user.Get("/get", userHandler.GetClerkUser)

	user.Post("/register", userHandler.RegisterUser)
	user.Post("/login", userHandler.Login)
	user.Put("/update", userHandler.UpdateUser)

	// Product Routes
	product := api.Group("/product")
	product.Get("/:id", productHandler.GetProductByID)
	product.Get("/", productHandler.ListProducts)
	product.Post("/", productHandler.CreateProduct)
	product.Put("/:id", productHandler.UpdateProduct)
	product.Delete("/:id", productHandler.DeleteProduct)

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

	// Order Routes
	// Not Needed for nowu
	// orderRoutes := app.Group("/orders")

	// orderRoutes.Post("/", orderHandler.PlaceOrder)
	// orderRoutes.Get("/:order_id", orderHandler.GetOrder)
	// orderRoutes.Get("/user/:user_id", orderHandler.GetUserOrders)
	// orderRoutes.Put("/:order_id", orderHandler.UpdateOrderStatus)
	// orderRoutes.Delete("/:order_id", orderHandler.CancelOrder)

	// Payment Routes
	// Not Needed for now
	// app.Post("/payment/create", paymentHandler.CreateOrder)
	// app.Get("/payment/verify/:transaction_id/:order_id", paymentHandler.VerifyPayment)

}
