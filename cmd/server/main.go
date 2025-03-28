package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/db"
	"github.com/shivajee98/aamishrit/internal/handlers"
	"github.com/shivajee98/aamishrit/internal/routes"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

func main() {
	app := fiber.New()

	// Connect to DB
	dbConn, err := db.Connect()

	utils.CheckError("Database Connection Failed!", err)

	// Initialising Handler with DB
	handler := handlers.Provide(dbConn)

	// Routes
	routes.Setup(app, handler)

	log.Fatal(app.Listen(":3000"))
}
