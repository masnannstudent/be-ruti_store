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

func (r *ProductRepository) CreateProduct(newData *entities.ProductModels, categoryIDs []uint64) (*entities.ProductModels, error) {

	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}

	if len(categoryIDs) > 0 {
		categories := make([]entities.CategoryModels, len(categoryIDs))
		for i, categoryID := range categoryIDs {
			categories[i] = entities.CategoryModels{ID: categoryID}
		}

		if err := r.db.Model(newData).Association("Categories").Append(categories); err != nil {
			return nil, err
		}
	}

	return newData, nil
}

func (r *ProductRepository) UpdateProduct(productID uint64, newData *entities.ProductModels, categoryIDs []uint64) error {
	var existingProduct *entities.ProductModels
	if err := r.db.Where("id = ?", productID).First(&existingProduct).Error; err != nil {
		return err
	}

	if err := r.db.Model(&existingProduct).Updates(newData).Error; err != nil {
		return err
	}

	if len(existingProduct.Categories) > 0 {
		if err := r.db.Model(existingProduct).Association("Categories").Delete(existingProduct.Categories); err != nil {
			return err
		}
	}

	if len(categoryIDs) > 0 {
		categories := make([]entities.CategoryModels, len(categoryIDs))
		for i, categoryID := range categoryIDs {
			categories[i] = entities.CategoryModels{ID: categoryID}
		}

		if err := r.db.Model(existingProduct).Association("Categories").Replace(categories); err != nil {
			return err
		}
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(productID uint64) error {
	var existingProduct *entities.ProductModels
	if err := r.db.Where("id = ?", productID).Preload("Categories").First(&existingProduct).Error; err != nil {
		return err
	}

	if len(existingProduct.Categories) > 0 {
		if err := r.db.Model(existingProduct).Association("Categories").Delete(&existingProduct.Categories); err != nil {
			return err
		}
	}

	if err := r.db.Delete(existingProduct).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) UpdateTotalReview(productID uint64) error {
	var products *entities.ProductModels
	err := r.db.Model(&products).Where("id = ?", productID).UpdateColumn("total_reviews", gorm.Expr("total_reviews + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) UpdateProductRating(productID uint64, newRating float64) error {
	if err := r.db.Model(&entities.ProductModels{}).Where("id = ?", productID).Update("rating", newRating).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetProductReviews(page, perPage int) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels
	offset := (page - 1) * perPage
	err := r.db.Where("deleted_at IS NULL").Offset(offset).Limit(perPage).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
