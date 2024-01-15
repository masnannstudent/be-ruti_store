package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/middleware"
	"ruti-store/module/feature/user/domain"
	"ruti-store/module/feature/user/handler"
	"ruti-store/module/feature/user/repository"
	"ruti-store/module/feature/user/service"
	"ruti-store/utils/token"
)

var (
	repo domain.UserRepositoryInterface
	serv domain.UserServiceInterface
	hand domain.UserHandlerInterface
)

func InitializeUser(db *gorm.DB) {
	repo = repository.NewUserRepository(db)
	serv = service.NewUserService(repo)
	hand = handler.NewUserHandler(serv)
}

func SetupRoutesUser(app *fiber.App, jwt token.JWTInterface, userService domain.UserServiceInterface) {
	api := app.Group("/api/v1/user")
	api.Get("/:id", middleware.AuthMiddleware(jwt, userService), hand.GetUserByID)
	api.Post("/get-profile", middleware.AuthMiddleware(jwt, userService), hand.GetUserProfile)
}
