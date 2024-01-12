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
	"ruti-store/utils/token"
)

var (
	repo domain.AddressRepositoryInterface
	serv domain.AddressServiceInterface
	hand domain.AddressHandlerInterface
)

func InitializeAddress(db *gorm.DB) {
	repo = repository.NewAddressRepository(db)
	serv = service.NewAddressService(repo)
	hand = handler.NewAddressHandler(serv)
}

func SetupRoutesAddress(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/address")
	api.Get("", middleware.AuthMiddleware(jwt, userService), hand.GetAllAddresses)
	api.Get("/:id", middleware.AuthMiddleware(jwt, userService), hand.GetAddressByID)
}
