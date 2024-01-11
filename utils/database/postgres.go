package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"ruti-store/config"
)

func InitPGSDatabase(config config.Config) *gorm.DB {
	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.Raw("SELECT 1").Error
	if err != nil {
		panic("failed to ping database")
	}

	fmt.Println("Connected to PostgreSQL!")
	return db
}
