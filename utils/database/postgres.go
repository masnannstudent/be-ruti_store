package database

import (
	"debtomate/utils/viper"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPGSDatabase() (*gorm.DB, error) {
	connection := viper.ViperConfig.GetStringValue("database.DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Check the connection
	if err := db.Exec("SELECT 1").Error; err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Connected to PostgresSQL!")
	return db, nil
}
