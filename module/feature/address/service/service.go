package service

import (
	"math"
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
