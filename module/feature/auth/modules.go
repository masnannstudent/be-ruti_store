package auth

import (
	"debtomate/module/feature/auth/domain"
	"debtomate/module/feature/auth/handler"
	"debtomate/module/feature/auth/repository"
	"debtomate/module/feature/auth/service"
	utils "debtomate/utils/hash"
	"debtomate/utils/token"
	"debtomate/utils/viper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	userRepo    domain.AuthRepositoryInterface
	userService domain.AuthServiceInterface
	userHandler domain.AuthHandlerInterface
	hash        utils.HashInterface
	jwt         token.JWTInterface
)

func InitializeAuth(db *gorm.DB) {
	secret := viper.ViperConfig.GetStringValue("app.SECRET")
	hash = utils.NewHash()
	jwt = token.NewJWT(secret)

	userRepo = repository.NewAuthRepository(db)
	userService = service.NewAuthService(userRepo, hash, jwt)
	userHandler = handler.NewAuthHandler(userService)
}

func SetupRoutesAuth(app *fiber.App) {
	api := app.Group("/api/v1/auth")
	api.Post("/login", userHandler.Login)
}
