package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/address/domain"
	"ruti-store/utils/shipping"
)

type AddressRepository struct {
	db       *gorm.DB
	shipping shipping.ShippingServiceInterface
}

func NewAddressRepository(db *gorm.DB, shipping shipping.ShippingServiceInterface) domain.AddressRepositoryInterface {
	return &AddressRepository{
		db:       db,
		shipping: shipping,
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

func (r *AddressRepository) CreateAddress(newData *entities.AddressModels) (*entities.AddressModels, error) {
	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *AddressRepository) GetPrimaryAddressByUserID(userID uint64) (*entities.AddressModels, error) {
	var addresses *entities.AddressModels
	err := r.db.Where("user_id = ? AND is_primary = ? AND deleted_at IS NULL", userID, true).First(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (r *AddressRepository) UpdateIsPrimary(addressID uint64, isPrimary bool) error {
	var addresses *entities.AddressModels
	err := r.db.Model(&addresses).Where("id = ?", addressID).Update("is_primary", isPrimary).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *AddressRepository) GetProvince() (map[string]interface{}, error) {
	result, err := r.shipping.GetProvince()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *AddressRepository) GetCity(province string) (map[string]interface{}, error) {
	result, err := r.shipping.GetCity(province)
	if err != nil {
		return nil, err
	}
	return result, nil
}
