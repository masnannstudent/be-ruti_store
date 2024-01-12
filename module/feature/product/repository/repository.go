package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/product/domain"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Model(&entities.ProductModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *ProductRepository) GetPaginatedProducts(page, pageSize int) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels

	offset := (page - 1) * pageSize

	if err := r.db.Offset(offset).Limit(pageSize).Preload("Photos").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
