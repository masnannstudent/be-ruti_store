package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/category/domain"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepositoryInterface {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.CategoryModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *CategoryRepository) GetPaginatedCategories(page, pageSize int) ([]*entities.CategoryModels, error) {
	var categories []*entities.CategoryModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).Limit(pageSize).Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
