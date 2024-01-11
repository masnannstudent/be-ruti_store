package route

import (
	"debtomate/module/feature/auth"
	"debtomate/utils/token"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwt token.JWTInterface) {
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
}
