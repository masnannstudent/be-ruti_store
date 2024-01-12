package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/address/domain"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) domain.AddressRepositoryInterface {
	return &AddressRepository{
		db: db,
	}
}

func (r *AddressRepository) GetAddressByID(addressID uint64) (*entities.AddressModels, error) {
	var address *entities.AddressModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", addressID).First(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *AddressRepository) GetTotalItems(userID uint64) (int64, error) {
	var totalItems int64

	if err := r.db.Model(&entities.AddressModels{}).Where("user_id = ? AND deleted_at IS NULL", userID).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *AddressRepository) GetPaginatedAddresses(userID uint64, page, pageSize int) ([]*entities.AddressModels, error) {
	var addresses []*entities.AddressModels
	offset := (page - 1) * pageSize
	err := r.db.Where("user_id = ? AND deleted_at IS NULL", userID).
		Limit(pageSize).
		Offset(offset).
		Find(&addresses).
		Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
