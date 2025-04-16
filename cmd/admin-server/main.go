package main

import (
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/db"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/repository"
	routes "github.com/shivajee98/aamishrit/internal/routes/admin-routes"
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
	categoryRepo := repository.InitCategoryRepository(dbConn)

	productService := services.InitProductService(productRepo)
	userService := services.InitUserService(userRepo)
	categoryService := services.InitCategoryService(categoryRepo)
	orderService := services.NewOrderService(orderRepo, userRepo)

	deps := routes.AdminDeps{
		ProductHandler:  handlers.InitProductHandler(productService, cloudinaryUploader),
		UserHandler:     handlers.InitUserHandler(userService),
		OrderHandler:    handlers.NewOrderHandler(orderService),
		CategoryHandler: handlers.InitCategoryHandler(categoryService, cloudinaryUploader),
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // allow admin panel from anywhere (or restrict as needed)
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// Mount only admin routes
	routes.SetupAdminRoutes(app, deps)

	port := os.Getenv("APORT")
	if port == "" {
		port = "3002" // or "3002"
	}
	log.Fatal(app.Listen(":" + port))
}
