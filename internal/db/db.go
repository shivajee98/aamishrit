package db

import (
	"github.com/shivajee98/aamishrit/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(db_uri string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(db_uri), &gorm.Config{})

	utils.CheckError("Error connecting to database", err)

	return db, nil
}
