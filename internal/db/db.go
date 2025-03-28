package db

import (
	"fmt"
	"os"

	"github.com/shivajee98/aamishrit/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	fmt.Println("Connecting to Database ....")

	db_uri := os.Getenv("SUPABASE_URL")

	db, err := gorm.Open(postgres.Open(db_uri))

	utils.CheckError("Error connecting to database", err)

	return db, nil
}
