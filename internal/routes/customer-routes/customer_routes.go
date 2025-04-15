package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/middleware"
)

func SetupCustomerRoutes(app *fiber.App, deps Deps) {
	cfg := config.LoadEnv()
	clerkKey := cfg.ClerkSecretKey

	api := app.Group("/api")

	// Public
	api.Post("/register", deps.UserHandler.RegisterUser)
	api.Post("/login", deps.UserHandler.Login)
	api.Get("/products/:id", deps.ProductHandler.GetProductByID)
	api.Get("/products", deps.ProductHandler.ListProducts)

	// Protected
	protected := api.Group("/", middleware.ClerkMiddleware(clerkKey))

	// User
	user := protected.Group("/user")
	user.Put("/", deps.UserHandler.UpdateUser)

	// Product (Admin/Seller maybe later)
	product := protected.Group("/products")
	product.Post("/", deps.ProductHandler.CreateProduct)
	product.Put("/:id", deps.ProductHandler.UpdateProduct)
	product.Delete("/:id", deps.ProductHandler.DeleteProduct)

	// Cart
	cart := protected.Group("/cart")
	cart.Post("/", deps.CartHandler.AddToCart)
	cart.Get("/", deps.CartHandler.GetCart)
	cart.Delete("/:cart_id", deps.CartHandler.RemoveFromCart)
	cart.Delete("/clear", deps.CartHandler.ClearCart)

	// Reviews
	review := protected.Group("/reviews")
	review.Post("/", deps.ReviewHandler.AddReview)
	review.Get("/:product_id", deps.ReviewHandler.GetReviews)
	review.Put("/:review_id", deps.ReviewHandler.UpdateReview)
	review.Delete("/:review_id", deps.ReviewHandler.DeleteReview)

	// Address
	address := protected.Group("/address")
	address.Get("/", deps.AddressHandler.GetAllAddresses)
	address.Post("/", deps.AddressHandler.CreateAddress)
	address.Get("/:id", deps.AddressHandler.GetAddressByID)
	address.Put("/:id", deps.AddressHandler.UpdateAddress)
	address.Delete("/:id", deps.AddressHandler.DeleteAddress)
	address.Put("/:id/default", deps.AddressHandler.SetDefaultAddress)
	address.Get("/default", deps.AddressHandler.GetDefaultAddress)

	// Orders
	orderRoutes := app.Group("/orders")
	orderRoutes.Post("/", deps.OrderHandler.PlaceOrder)
	orderRoutes.Get("/:order_id", deps.OrderHandler.GetOrder)
	orderRoutes.Get("/user/:user_id", deps.OrderHandler.GetUserOrders)
	orderRoutes.Put("/:order_id", deps.OrderHandler.UpdateOrderStatus)
	orderRoutes.Delete("/:order_id", deps.OrderHandler.CancelOrder)

	// category
	category := app.Group("/categories")
	category.Get("/", deps.CategoryHandler.GetCategories)
}
