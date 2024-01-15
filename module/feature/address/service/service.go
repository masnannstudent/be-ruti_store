package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/address/domain"
	"time"
)

type AddressService struct {
	repo domain.AddressRepositoryInterface
}

func NewAddressService(repo domain.AddressRepositoryInterface) domain.AddressServiceInterface {
	return &AddressService{
		repo: repo,
	}
}

func (s *AddressService) GetAddressByID(addressID uint64) (*entities.AddressModels, error) {
	result, err := s.repo.GetAddressByID(addressID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *AddressService) GetAllAddresses(userID uint64, page, pageSize int) ([]*entities.AddressModels, int64, error) {
	result, err := s.repo.GetPaginatedAddresses(userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems(userID)
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *AddressService) GetAddressesPage(userID uint64, currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalItems(userID)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return currentPage, totalPages, nextPage, prevPage, nil
}

func (s *AddressService) CreateAddress(userID uint64, req *domain.CreateAddressRequest) (*entities.AddressModels, error) {
	newData := &entities.AddressModels{
		UserID:       userID,
		AcceptedName: req.AcceptedName,
		Phone:        req.Phone,
		ProvinceID:   req.ProvinceID,
		ProvinceName: req.ProvinceName,
		CityID:       req.CityID,
		CityName:     req.CityName,
		Address:      req.Address,
		IsPrimary:    req.IsPrimary,
		CreatedAt:    time.Now(),
	}

	createdAddress, err := s.repo.CreateAddress(newData)
	if err != nil {
		return nil, err
	}
	if createdAddress.IsPrimary {
		currentPrimaryAddress, err := s.repo.GetPrimaryAddressByUserID(createdAddress.UserID)
		if err != nil {
			return nil, errors.New("gagal mendapatkan alamat utama")
		}
		if currentPrimaryAddress != nil && currentPrimaryAddress.ID != createdAddress.ID {
			err = s.repo.UpdateIsPrimary(currentPrimaryAddress.ID, false)
			if err != nil {
				return nil, errors.New("gagal merubah alamat utama")
			}
		}
	}

	return createdAddress, nil
}

func (s *AddressService) GetProvince() (map[string]interface{}, error) {
	result, err := s.repo.GetProvince()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *AddressService) GetCity(province string) (map[string]interface{}, error) {
	result, err := s.repo.GetCity(province)
	if err != nil {
		return nil, err
	}
	return result, nil
}
