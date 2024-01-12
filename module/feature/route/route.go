package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/auth"
	"ruti-store/module/feature/product"
	"ruti-store/utils/token"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwt token.JWTInterface) {
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
	product.InitializeAuth(db)
	product.SetupRoutesAuth(app)
}
