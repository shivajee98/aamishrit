package main

import (
	"log"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/db"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/middleware"
	"github.com/shivajee98/aamishrit/internal/routes"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

func main() {
	app := fiber.New()

	// Connect to DB
	dbConn, err := db.Connect()

	// Load Configs
	cfg := config.LoadConfig()

	// Initialising Clerk and Storing JWK

	clerkConfig := &clerk.ClientConfig{}
	clerkConfig.Key = clerk.String(cfg.ClerkSecretKey)
	jwkClient := jwks.NewClient(clerkConfig)
	jwkStore := middleware.NewJWKStore()

	// Initialising Handler with DB
	handler := handlers.Provide(dbConn)

	// Register Routes
	routes.Setup(app, jwkClient, jwkStore, handler)

	utils.CheckError("Database Connection Failed!", err)

	log.Fatal(app.Listen(":3000"))
}
