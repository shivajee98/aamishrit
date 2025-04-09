package main

import (
	"log"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/db"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/repository"
	"github.com/shivajee98/aamishrit/internal/routes"
	"github.com/shivajee98/aamishrit/internal/services"
	"github.com/shivajee98/aamishrit/internal/uploader"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

func main() {
	app := fiber.New()

	// Load Configs
	cfg := config.LoadEnv()

	// Connect to DB
	dbConn, err := db.Connect(cfg.DatabaseURL)

	// clerk key setup
	clerk.SetKey(cfg.CLERK_SECRET_KEY)

	// cloudinary
	cloudinaryUploader := uploader.NewCloudinaryUploader(cfg)

	// User
	setupuserRepo := repository.InitUserRepository(dbConn)
	userService := services.InitUserService(setupuserRepo)
	userHandler := handlers.InitUserHandler(userService)

	// product
	productRepo := repository.InitProductRepository(dbConn)
	productService := services.InitProductService(productRepo)
	productHandler := handlers.InitProductHandler(productService, cloudinaryUploader)

	// cart
	cartRepo := repository.InitCartRepository(dbConn)
	cartService := services.InitCartService(cartRepo)
	cartHandler := handlers.InitCartHandler(cartService)

	// Review
	reviewRepo := repository.InitReviewRepository(dbConn)
	reviewService := services.InitReviewService(reviewRepo)
	reviewHandler := handlers.InitReviewHandler(reviewService)

	// Address
	addressRepo := repository.InitAddressRepository(dbConn)
	addressService := services.InitAddressService(addressRepo)
	addressHandler := handlers.InitAddressHandler(addressService)

	// Order
	orderRepo := repository.NewOrderRepository(dbConn)
	orderService := services.NewOrderService(orderRepo, setupuserRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Payment
	// Not Needed for now
	// paymentRepo := repository.NewPaymentRepository(dbConn)
	// paymentService := services.NewPaymentService(paymentRepo)
	// paymentHandler := handlers.NewPaymentHandler(paymentService)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",      // Allow the specific origin
		AllowMethods: "GET,POST,PUT,DELETE",        // Methods that are allowed
		AllowHeaders: "Content-Type,Authorization", // Headers that are allowed
	}))

	routes.Setup(app, userHandler, productHandler, cartHandler, reviewHandler, addressHandler, orderHandler)

	utils.CheckError("Database Connection Failed!", err)

	log.Fatal(app.Listen(":3000"))
}
