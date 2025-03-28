package migrate

import (
	"fmt"
	"log"

	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/internal/db"
	model "github.com/shivajee98/aamishrit/internal/models"
)

func main() {
	fmt.Println("Starting migration...")

	config.LoadEnv()

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = dbConn.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}
