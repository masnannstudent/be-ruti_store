package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/category/domain"
	"time"
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

func (r *CategoryRepository) GetCategoryByID(categoryID uint64) (*entities.CategoryModels, error) {
	var category *entities.CategoryModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", categoryID).Preload("Products").First(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) CreateCategory(category *entities.CategoryModels) (*entities.CategoryModels, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) UpdateCategory(categoryID uint64, updatedCategory *entities.CategoryModels) error {
	var category *entities.CategoryModels
	if err := r.db.Where("id = ? AND deleted_at IS NULL", categoryID).First(&category).Error; err != nil {
		return err
	}

	if err := r.db.Model(category).Updates(updatedCategory).Error; err != nil {
		return err
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(categoryID uint64) error {
	category := &entities.CategoryModels{}
	if err := r.db.First(category, categoryID).Error; err != nil {
		return err
	}

	if err := r.db.Model(category).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
