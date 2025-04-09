package main

import (
	"fmt"
	"log"

	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/db"
	"github.com/shivajee98/aamishrit/internal/model"
)

func main() {
	fmt.Println("Starting migration...")

	cfg := config.LoadEnv()

	dbConn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Drop the tables if they exist
	err = dbConn.Migrator().DropTable(&model.User{}, &model.Product{})
	if err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	err = dbConn.AutoMigrate(&model.User{}, &model.Product{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}
