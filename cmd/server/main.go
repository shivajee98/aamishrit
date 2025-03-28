package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivajee98/aamishrit/internal/db"
	"github.com/shivajee98/aamishrit/internal/routes"
	"github.com/shivajee98/aamishrit/pkg/utils"
)

func main() {
	app := fiber.New()

	// Connect to DB
	dbConn, err := db.Connect()

	utils.CheckError("Database Connection Failed!", err)

	// Routes
	routes.Setup(app, dbConn)
}
