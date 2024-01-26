package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"ruti-store/module/feature/address/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"ruti-store/module/entities"
	"ruti-store/module/feature/address/mocks"
)

func cloneAddress(a *entities.AddressModels) *entities.AddressModels {
	return &entities.AddressModels{
		ID:           a.ID,
		UserID:       a.UserID,
		AcceptedName: a.AcceptedName,
		Phone:        a.Phone,
		ProvinceID:   a.ProvinceID,
		ProvinceName: a.ProvinceName,
		CityID:       a.CityID,
		CityName:     a.CityName,
		Address:      a.Address,
		IsPrimary:    a.IsPrimary,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
		DeletedAt:    a.DeletedAt,
	}
}

func TestAddressService_GetAddressByID(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	address := &entities.AddressModels{
		ID:           1,
		UserID:       123,
		AcceptedName: "John Doe",
		Phone:        "123456789",
		ProvinceID:   "P01",
		ProvinceName: "Province 1",
		CityID:       "C01",
		CityName:     "City 1",
		Address:      "Street 123",
		IsPrimary:    true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}

	expectedAddress := cloneAddress(address)

	t.Run("Success Case - Address Found", func(t *testing.T) {
		addressID := uint64(1)
		repo.On("GetAddressByID", addressID).Return(expectedAddress, nil).Once()

		result, err := service.GetAddressByID(addressID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedAddress.ID, result.ID)
		assert.Equal(t, expectedAddress.UserID, result.UserID)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Address Not Found", func(t *testing.T) {
		addressID := uint64(2)

		expectedErr := errors.New("address not found")
		repo.On("GetAddressByID", addressID).Return(nil, expectedErr).Once()

		result, err := service.GetAddressByID(addressID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestAddressService_GetAllAddresses(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	userID := uint64(123)
	page := 1
	pageSize := 10

	addresses := []*entities.AddressModels{
		{
			ID:           1,
			UserID:       userID,
			AcceptedName: "John Doe",
			Phone:        "123456789",
			ProvinceID:   "P01",
			ProvinceName: "Province 1",
			CityID:       "C01",
			CityName:     "City 1",
			Address:      "Street 123",
			IsPrimary:    true,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    nil,
		},
	}

	expectedAddresses := []*entities.AddressModels{cloneAddress(addresses[0])}

	t.Run("Success Case - Addresses Found", func(t *testing.T) {
		repo.On("GetPaginatedAddresses", userID, page, pageSize).Return(expectedAddresses, nil).Once()
		repo.On("GetTotalItems", userID).Return(int64(len(addresses)), nil).Once()

		result, totalItems, err := service.GetAllAddresses(userID, page, pageSize)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(expectedAddresses), len(result))
		assert.Equal(t, len(addresses), int(totalItems))

		for i := range result {
			assert.Equal(t, expectedAddresses[i].ID, result[i].ID)
			assert.Equal(t, expectedAddresses[i].UserID, result[i].UserID)
		}

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Getting Addresses", func(t *testing.T) {
		expectedErr := errors.New("error getting addresses")
		repo.On("GetPaginatedAddresses", userID, page, pageSize).Return(nil, expectedErr).Once()

		result, _, err := service.GetAllAddresses(userID, page, pageSize)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Getting Total Items", func(t *testing.T) {
		expectedErr := errors.New("error getting total items")
		repo.On("GetPaginatedAddresses", userID, page, pageSize).Return(expectedAddresses, nil).Once()
		repo.On("GetTotalItems", userID).Return(int64(0), expectedErr).Once()

		result, _, err := service.GetAllAddresses(userID, page, pageSize)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestAddressService_GetAddressesPage(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	userID := uint64(123)
	pageSize := 10

	t.Run("Success Case", func(t *testing.T) {
		currentPage := 2
		expectedTotalItems := int64(25)
		repo.On("GetTotalItems", userID).Return(expectedTotalItems, nil).Once()

		expectedTotalPages := 3
		expectedNextPage := currentPage + 1
		expectedPrevPage := currentPage - 1

		if expectedNextPage > expectedTotalPages {
			expectedNextPage = 0
		}

		if expectedPrevPage < 1 {
			expectedPrevPage = 0
		}

		resultCurrentPage, resultTotalPages, resultNextPage, resultPrevPage, err := service.GetAddressesPage(userID, currentPage, pageSize)

		assert.Nil(t, err)
		assert.Equal(t, currentPage, resultCurrentPage)
		assert.Equal(t, expectedTotalPages, resultTotalPages)
		assert.Equal(t, expectedNextPage, resultNextPage)
		assert.Equal(t, expectedPrevPage, resultPrevPage)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Getting Total Items", func(t *testing.T) {
		currentPage := 2
		expectedErr := errors.New("error getting total items")
		repo.On("GetTotalItems", userID).Return(int64(0), expectedErr).Once()

		resultCurrentPage, resultTotalPages, resultNextPage, resultPrevPage, err := service.GetAddressesPage(userID, currentPage, pageSize)

		assert.Error(t, err)
		assert.Zero(t, resultCurrentPage)
		assert.Zero(t, resultTotalPages)
		assert.Zero(t, resultNextPage)
		assert.Zero(t, resultPrevPage)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestAddressService_CreateAddress(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	userID := uint64(123)
	createAddressRequest := &domain.CreateAddressRequest{
		AcceptedName: "John Doe",
		Phone:        "123456789",
		ProvinceID:   "P01",
		ProvinceName: "Province 1",
		CityID:       "C01",
		CityName:     "City 1",
		Address:      "Street 123",
		IsPrimary:    true,
	}

	t.Run("Failed Case - Error Creating Address", func(t *testing.T) {
		expectedErr := errors.New("error creating address")
		repo.On("CreateAddress", mock.Anything).Return(nil, expectedErr).Once()

		result, err := service.CreateAddress(userID, createAddressRequest)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestAddressService_GetProvince(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	t.Run("Success Case - Get Province", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"province1": "Province 1",
			"province2": "Province 2",
		}

		repo.On("GetProvince").Return(expectedResult, nil).Once()

		result, err := service.GetProvince()

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Getting Province", func(t *testing.T) {
		expectedErr := errors.New("error getting province")
		repo.On("GetProvince").Return(nil, expectedErr).Once()

		result, err := service.GetProvince()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestAddressService_GetCity(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	province := "Province1"

	t.Run("Success Case - Get City", func(t *testing.T) {
		expectedResult := map[string]interface{}{
			"city1": "City 1",
			"city2": "City 2",
		}

		repo.On("GetCity", province).Return(expectedResult, nil).Once()

		result, err := service.GetCity(province)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Getting City", func(t *testing.T) {
		expectedErr := errors.New("error getting city")
		repo.On("GetCity", province).Return(nil, expectedErr).Once()

		result, err := service.GetCity(province)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})
}

func TestAddressService_UpdateAddress(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	addressID := uint64(1)
	updateAddressRequest := &domain.UpdateAddressRequest{
		AcceptedName: "John Doe",
		Phone:        "123456789",
		ProvinceID:   "P01",
		ProvinceName: "Province 1",
		CityID:       "C01",
		CityName:     "City 1",
		Address:      "Street 123",
		IsPrimary:    true,
	}

	t.Run("Failed Case - Error Getting Address", func(t *testing.T) {
		expectedErr := errors.New("address not found")
		repo.On("GetAddressByID", addressID).Return(nil, expectedErr).Once()

		result, err := service.UpdateAddress(addressID, updateAddressRequest)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})

}

func TestAddressService_DeleteAddress(t *testing.T) {
	repo := mocks.NewAddressRepositoryInterface(t)
	service := NewAddressService(repo)

	addressID := uint64(1)
	userID := uint64(123)

	t.Run("Failed Case - Error Getting Address", func(t *testing.T) {
		expectedErr := errors.New("address not found")
		repo.On("GetAddressByID", addressID).Return(nil, expectedErr).Once()

		err := service.DeleteAddress(addressID, userID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Getting Primary Address", func(t *testing.T) {
		address := &entities.AddressModels{
			ID:           addressID,
			UserID:       userID,
			AcceptedName: "John Doe",
			Phone:        "123456789",
			ProvinceID:   "P01",
			ProvinceName: "Province 1",
			CityID:       "C01",
			CityName:     "City 1",
			Address:      "Street 123",
			IsPrimary:    false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		expectedErr := errors.New("failed to get primary address")
		repo.On("GetAddressByID", addressID).Return(address, nil).Once()
		repo.On("GetPrimaryAddressByUserID", userID).Return(nil, expectedErr).Once()

		err := service.DeleteAddress(addressID, userID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)

		repo.AssertExpectations(t)
	})

	t.Run("Success Case - Delete Address", func(t *testing.T) {
		address := &entities.AddressModels{
			ID:           addressID,
			UserID:       userID,
			AcceptedName: "John Doe",
			Phone:        "123456789",
			ProvinceID:   "P01",
			ProvinceName: "Province 1",
			CityID:       "C01",
			CityName:     "City 1",
			Address:      "Street 123",
			IsPrimary:    false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		expectedPrimaryAddress := &entities.AddressModels{
			ID:           2,
			UserID:       userID,
			AcceptedName: "Jane Doe",
			Phone:        "987654321",
			ProvinceID:   "P02",
			ProvinceName: "Province 2",
			CityID:       "C02",
			CityName:     "City 2",
			Address:      "Street 456",
			IsPrimary:    false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		repo.On("GetAddressByID", addressID).Return(address, nil).Once()
		repo.On("GetPrimaryAddressByUserID", userID).Return(expectedPrimaryAddress, nil).Once()
		repo.On("DeleteAddress", addressID).Return(nil).Once()

		err := service.DeleteAddress(addressID, userID)

		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})
}
