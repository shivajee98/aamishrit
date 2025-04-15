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
	"github.com/shivajee98/aamishrit/internal/routes/admin-routes"
	"github.com/shivajee98/aamishrit/internal/services"
	"github.com/shivajee98/aamishrit/internal/uploader"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	})
	cfg := config.LoadEnv()

	utils.CheckError("Loading Config Failed", nil)

	clerk.SetKey(cfg.CLERK_SECRET_KEY)

	dbConn, err := db.Connect(cfg.DatabaseURL)
	utils.CheckError("DB connection failed", err)

	cloudinaryUploader := uploader.NewCloudinaryUploader(cfg)

	// Wire Admin-only Repos, Services, Handlers
	productRepo := repository.InitProductRepository(dbConn)
	userRepo := repository.InitUserRepository(dbConn)
	orderRepo := repository.NewOrderRepository(dbConn)

	productService := services.InitProductService(productRepo)
	userService := services.InitUserService(userRepo)
	orderService := services.NewOrderService(orderRepo, userRepo)

	deps := routes.AdminDeps{
		ProductHandler: handlers.InitProductHandler(productService, cloudinaryUploader),
		UserHandler:    handlers.InitUserHandler(userService),
		OrderHandler:   handlers.NewOrderHandler(orderService),
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // allow admin panel from anywhere (or restrict as needed)
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// Mount only admin routes
	routes.SetupAdminRoutes(app, deps)

	log.Fatal(app.Listen(":3002")) // different port from customer
}
