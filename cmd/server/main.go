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
	productRepo := repository.InitProductRepository(dbConn)
	productService := services.InitProductService(productRepo)
	productHandler := handlers.InitProductHandler(productService)

	routes.Setup(app, productHandler)

	utils.CheckError("Database Connection Failed!", err)

	log.Fatal(app.Listen(":3000"))
}
