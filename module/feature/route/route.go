package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
	"ruti-store/module/feature/address"
	"ruti-store/module/feature/auth"
	"ruti-store/module/feature/home"
	"ruti-store/module/feature/order"
	"ruti-store/module/feature/product"
	users "ruti-store/module/feature/user"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/token"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwt token.JWTInterface,
	snapClient snap.Client, userService user.UserServiceInterface, coreClient coreapi.Client) {
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
	product.InitializeProduct(db)
	product.SetupRoutesProduct(app)
	order.InitializeOrder(db, snapClient, coreClient)
	order.SetupOrderRoutes(app, jwt, userService)
	address.InitializeAddress(db)
	address.SetupRoutesAddress(app, jwt, userService)
	home.InitializeHome(db)
	home.SetupRoutesHome(app, jwt, userService)
	users.InitializeUser(db)
	users.SetupRoutesUser(app, jwt, userService)
}
