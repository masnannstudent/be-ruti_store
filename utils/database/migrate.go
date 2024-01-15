package database

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entities.UserModels{},
		entities.AddressModels{},
		entities.ProductModels{},
		entities.ProductPhotoModels{},
		entities.CategoryModels{},
		entities.OrderModels{},
		entities.OrderDetailsModels{},
		entities.CarouselModels{})

	if err != nil {
		return
	}
	return
}
