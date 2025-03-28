package db

import (
	"github.com/shivajee98/aamishrit/internal/config"
	"github.com/shivajee98/aamishrit/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	cfg := config.LoadEnv()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})

	utils.CheckError("Error connecting to database", err)

	return db, nil
}
