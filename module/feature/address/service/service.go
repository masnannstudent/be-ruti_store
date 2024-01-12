package service

import (
	"ruti-store/module/entities"
	"ruti-store/module/feature/address/domain"
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
