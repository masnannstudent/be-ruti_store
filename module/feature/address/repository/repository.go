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
