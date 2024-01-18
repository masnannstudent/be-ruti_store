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
		entities.CarouselModels{},
		entities.ReviewModels{},
		entities.ReviewPhotoModels{},
		entities.ArticleModels{},
		entities.NotificationModels{},
		entities.CartModels{})

	if err != nil {
		return
	}
	return
}
