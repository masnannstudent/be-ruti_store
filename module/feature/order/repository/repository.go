package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/order/domain"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepositoryInterface {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Model(&entities.OrderModels{}).Count(&totalItems).Where("deleted_at IS NULL").Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *OrderRepository) GetPaginatedOrders(page, pageSize int) ([]*entities.OrderModels, error) {
	var orders []*entities.OrderModels

	offset := (page - 1) * pageSize

	err := r.db.Offset(offset).Limit(pageSize).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Preload("User").
		Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
