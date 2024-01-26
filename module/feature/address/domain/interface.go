package domain

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
)

type AddressRepositoryInterface interface {
	GetAddressByID(addressID uint64) (*entities.AddressModels, error)
	GetTotalItems(userID uint64) (int64, error)
	GetPaginatedAddresses(userID uint64, page, pageSize int) ([]*entities.AddressModels, error)
	CreateAddress(newData *entities.AddressModels) (*entities.AddressModels, error)
	GetPrimaryAddressByUserID(userID uint64) (*entities.AddressModels, error)
	UpdateIsPrimary(addressID uint64, isPrimary bool) error
	GetProvince() (map[string]interface{}, error)
	GetCity(province string) (map[string]interface{}, error)
	UpdateAddress(addressID uint64, updatedAddress *entities.AddressModels) (*entities.AddressModels, error)
	DeleteAddress(addressID uint64) error
}

type AddressServiceInterface interface {
	GetAddressByID(addressID uint64) (*entities.AddressModels, error)
	GetAllAddresses(userID uint64, page, pageSize int) ([]*entities.AddressModels, int64, error)
	GetAddressesPage(userID uint64, currentPage, pageSize int) (int, int, int, int, error)
	CreateAddress(userID uint64, req *CreateAddressRequest) (*entities.AddressModels, error)
	GetProvince() (map[string]interface{}, error)
	GetCity(province string) (map[string]interface{}, error)
	UpdateAddress(addressID uint64, req *UpdateAddressRequest) (*entities.AddressModels, error)
	DeleteAddress(addressID, userID uint64) error
}

type AddressHandlerInterface interface {
	GetAddressByID(c *fiber.Ctx) error
	GetAllAddresses(c *fiber.Ctx) error
	CreateAddress(c *fiber.Ctx) error
	GetProvince(c *fiber.Ctx) error
	GetCity(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	DeleteAddress(c *fiber.Ctx) error
}
