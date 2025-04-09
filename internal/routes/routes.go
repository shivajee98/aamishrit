package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/middleware"
)

func Setup(app *fiber.App, userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler, cartHandler *handlers.CartHandler, reviewHandler *handlers.ReviewHandler, addressHandler *handlers.AddressHandler) {
	cfg := config.LoadEnv()
	clerkKey := cfg.ClerkSecretKey

	api := app.Group("/api")

	// 🟢 Public Routes
	api.Post("/register", userHandler.RegisterUser)
	api.Post("/login", userHandler.Login)

	api.Get("/products/:id", productHandler.GetProductByID)
	api.Get("/products", productHandler.ListProducts)

	// 🔒 Protected Routes (Clerk Auth Middleware)
	protected := api.Group("/", middleware.ClerkMiddleware(clerkKey))

	// 🔐 User Routes
	user := protected.Group("/user")
	user.Put("/", userHandler.UpdateUser)

	// 🔐 Product Management (for admins / sellers)
	product := protected.Group("/products")
	product.Post("/", productHandler.CreateProduct)
	product.Put("/:id", productHandler.UpdateProduct)
	product.Delete("/:id", productHandler.DeleteProduct)

	// 🔐 Cart Routes
	cart := protected.Group("/cart")
	cart.Post("/", cartHandler.AddToCart)
	cart.Get("/", cartHandler.GetCart) // gets cart of current user
	cart.Delete("/:cart_id", cartHandler.RemoveFromCart)
	cart.Delete("/clear", cartHandler.ClearCart) // clears current user's cart

	// 🔐 Review Routes
	review := protected.Group("/reviews")
	review.Post("/", reviewHandler.AddReview)
	review.Get("/:product_id", reviewHandler.GetReviews)
	review.Put("/:review_id", reviewHandler.UpdateReview)
	review.Delete("/:review_id", reviewHandler.DeleteReview)

	// 🔐 Address Routes
	address := protected.Group("/address")
	address.Get("/", addressHandler.GetAllAddresses)
	address.Post("/", addressHandler.CreateAddress)
	address.Get("/:id", addressHandler.GetAddressByID)
	address.Put("/:id", addressHandler.UpdateAddress)
	address.Delete("/:id", addressHandler.DeleteAddress)
	address.Put("/:id/default", addressHandler.SetDefaultAddress)
	address.Get("/default", addressHandler.GetDefaultAddress)

	// Order Routes
	// Not Needed for now
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
