package entities

import "time"

type ProductModels struct {
	ID           uint64               `gorm:"column:id;primaryKey" json:"id"`
	Name         string               `gorm:"column:name" json:"name"`
	Price        uint64               `gorm:"column:price" json:"price"`
	Description  string               `gorm:"column:description" json:"description"`
	Discount     uint64               `gorm:"column:discount" json:"discount"`
	Rating       float64              `gorm:"column:rating" json:"rating"`
	TotalReviews uint64               `gorm:"column:total_reviews" json:"total_reviews"`
	Stock        uint64               `gorm:"column:stock" json:"stock"`
	CreatedAt    time.Time            `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt    time.Time            `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt    *time.Time           `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
	Photos       []ProductPhotoModels `gorm:"foreignKey:ProductID" json:"photos"`
	Categories   []*CategoryModels    `gorm:"many2many:product_categories;" json:"categories"`
}

type ProductPhotoModels struct {
	ID        uint64 `gorm:"column:id;primaryKey" json:"id"`
	ProductID uint64 `gorm:"column:product_id" json:"product_id"`
	URL       string `gorm:"column:url" json:"url"`
}

func (ProductModels) TableName() string {
	return "product"
}

func (ProductPhotoModels) TableName() string {
	return "product_photo"
}
