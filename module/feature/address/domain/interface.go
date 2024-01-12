package domain

import "ruti-store/module/entities"

type AddressRepositoryInterface interface {
	GetAddressByID(addressID uint64) (*entities.AddressModels, error)
}

type AddressServiceInterface interface {
	GetAddressByID(addressID uint64) (*entities.AddressModels, error)
}

type AddressHandlerInterface interface {
}