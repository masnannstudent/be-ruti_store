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

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.ProductModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *ProductRepository) GetPaginatedProducts(page, pageSize int) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).Limit(pageSize).Preload("Photos").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(productID uint64) (*entities.ProductModels, error) {
	var product *entities.ProductModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", productID).Preload("Photos").First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) CreateProduct(product *entities.ProductModels, categoryIDs []uint64) (*entities.ProductModels, error) {

	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}

	if len(categoryIDs) > 0 {
		categories := make([]entities.CategoryModels, len(categoryIDs))
		for i, categoryID := range categoryIDs {
			categories[i] = entities.CategoryModels{ID: categoryID}
		}

		if err := r.db.Model(product).Association("Categories").Replace(categories); err != nil {
			return nil, err
		}
	}

	return product, nil
}
