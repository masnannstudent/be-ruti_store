package address

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/address/domain"
	"ruti-store/module/feature/address/handler"
	"ruti-store/module/feature/address/repository"
	"ruti-store/module/feature/address/service"
	"ruti-store/module/feature/middleware"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/shipping"
	"ruti-store/utils/token"
)

var (
	repo domain.AddressRepositoryInterface
	serv domain.AddressServiceInterface
	hand domain.AddressHandlerInterface
	ship shipping.ShippingServiceInterface
)

func InitializeAddress(db *gorm.DB) {
	ship = shipping.NewShippingService()
	repo = repository.NewAddressRepository(db, ship)
	serv = service.NewAddressService(repo)
	hand = handler.NewAddressHandler(serv)
}

func SetupRoutesAddress(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/address")
	api.Get("", middleware.AuthMiddleware(jwt, userService), hand.GetAllAddresses)
	api.Get("/:id", middleware.AuthMiddleware(jwt, userService), hand.GetAddressByID)
	api.Post("", middleware.AuthMiddleware(jwt, userService), hand.CreateAddress)
	api.Post("/get-province", middleware.AuthMiddleware(jwt, userService), hand.GetProvince)
	api.Post("/get-city", middleware.AuthMiddleware(jwt, userService), hand.GetCity)
}
