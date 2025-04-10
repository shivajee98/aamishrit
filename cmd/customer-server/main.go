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
	routes "github.com/shivajee98/aamishrit/internal/routes/customer-routes"
	"github.com/shivajee98/aamishrit/internal/services"
	"github.com/shivajee98/aamishrit/internal/uploader"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

func main() {
	app := fiber.New()
	cfg := config.LoadEnv()
	utils.CheckError("Database Connection Failed!", nil)

	clerk.SetKey(cfg.CLERK_SECRET_KEY)

	dbConn, err := db.Connect(cfg.DatabaseURL)
	utils.CheckError("DB connection failed", err)

	cloudinaryUploader := uploader.NewCloudinaryUploader(cfg)

	// 🔌 Wire Dependencies
	userRepo := repository.InitUserRepository(dbConn)
	productRepo := repository.InitProductRepository(dbConn)
	cartRepo := repository.InitCartRepository(dbConn)
	reviewRepo := repository.InitReviewRepository(dbConn)
	addressRepo := repository.InitAddressRepository(dbConn)
	orderRepo := repository.NewOrderRepository(dbConn)

	userService := services.InitUserService(userRepo)
	productService := services.InitProductService(productRepo)
	cartService := services.InitCartService(cartRepo)
	reviewService := services.InitReviewService(reviewRepo)
	addressService := services.InitAddressService(addressRepo)
	orderService := services.NewOrderService(orderRepo, userRepo)

	deps := routes.Deps{
		UserHandler:    handlers.InitUserHandler(userService),
		ProductHandler: handlers.InitProductHandler(productService, cloudinaryUploader),
		CartHandler:    handlers.InitCartHandler(cartService),
		ReviewHandler:  handlers.InitReviewHandler(reviewService),
		AddressHandler: handlers.InitAddressHandler(addressService),
		OrderHandler:   handlers.NewOrderHandler(orderService),
	}

	// 🔒 Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// 🔀 Setup Routes
	routes.SetupCustomerRoutes(app, deps)

	log.Fatal(app.Listen(":3000"))
}
