package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type AddressRepositoryInterface interface {
	GetAddressByID(addressID uint64) (*entities.AddressModels, error)
	GetTotalItems(userID uint64) (int64, error)
	GetPaginatedAddresses(userID uint64, page, pageSize int) ([]*entities.AddressModels, error)
}

type AddressServiceInterface interface {
	GetAddressByID(addressID uint64) (*entities.AddressModels, error)
	GetAllAddresses(userID uint64, page, pageSize int) ([]*entities.AddressModels, int64, error)
	GetAddressesPage(userID uint64, currentPage, pageSize int) (int, int, int, int, error)
}

type AddressHandlerInterface interface {
	GetAddressByID(c *fiber.Ctx) error
	GetAllAddresses(c *fiber.Ctx) error
}
