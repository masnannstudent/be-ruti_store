package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
	"ruti-store/module/feature/auth/domain"
	"ruti-store/module/feature/auth/handler"
	"ruti-store/module/feature/auth/repository"
	"ruti-store/module/feature/auth/service"
	utils "ruti-store/utils/hash"
	"ruti-store/utils/token"
)

var (
	userRepo    domain.AuthRepositoryInterface
	userService domain.AuthServiceInterface
	userHandler domain.AuthHandlerInterface
	hash        utils.HashInterface
	jwt         token.JWTInterface
)

func InitializeAuth(db *gorm.DB) {
	secret := os.Getenv("SECRET")
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
