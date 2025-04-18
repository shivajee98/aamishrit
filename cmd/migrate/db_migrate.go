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

	// err = dbConn.Exec("TRUNCATE TABLE product_categories RESTART IDENTITY CASCADE").Error
	// if err != nil {
	// 	log.Fatalf("Failed to truncate product_categories: %v", err)
	// }

	// // Drop the tables if they exist
	// err = dbConn.Migrator().DropTable(&model.User{}, &model.Address{}, &model.Cart{}, &model.Order{}, &model.OrderItem{}, &model.Product{}, &model.Review{}, &model.Wishlist{}, &model.Category{})
	// if err != nil {
	// 	log.Fatalf("Failed to drop tables: %v", err)
	// }

	err = dbConn.AutoMigrate(&model.User{}, &model.Product{}, &model.Address{}, &model.Cart{}, &model.Order{}, &model.OrderItem{}, &model.Review{}, &model.Wishlist{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}
