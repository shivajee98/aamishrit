package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/db"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/repository"
	"github.com/shivajee98/aamishrit/internal/routes"
	"github.com/shivajee98/aamishrit/internal/services"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

func main() {
	app := fiber.New()

	// Load Configs
	cfg := config.LoadEnv()

	// Connect to DB
	dbConn, err := db.Connect(cfg.DatabaseURL)

	// setup
	// product
	productRepo := repository.InitProductRepository(dbConn)
	productService := services.InitProductService(productRepo)
	productHandler := handlers.InitProductHandler(productService)

	// cart
	cartRepo := repository.InitCartRepository(dbConn)
	cartService := services.InitCartService(cartRepo)
	cartHandler := handlers.InitCartHandler(cartService)

	// Review
	reviewRepo := repository.NewReviewRepository(dbConn)
	reviewService := services.NewReviewService(reviewRepo)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	// Order
	// Not Needed for now
	// orderRepo := repository.NewOrderRepository(dbConn)
	// orderService := services.NewOrderService(orderRepo)
	// orderHandler := handlers.NewOrderHandler(orderService)

	// Payment
	// Not Needed for now
	// paymentRepo := repository.NewPaymentRepository(dbConn)
	// paymentService := services.NewPaymentService(paymentRepo)
	// paymentHandler := handlers.NewPaymentHandler(paymentService)

	routes.Setup(app, productHandler, cartHandler, reviewHandler)

	utils.CheckError("Database Connection Failed!", err)

	log.Fatal(app.Listen(":3000"))
}
