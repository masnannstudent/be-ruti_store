package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/middleware"
	"ruti-store/module/feature/user/domain"
	"ruti-store/module/feature/user/handler"
	"ruti-store/module/feature/user/repository"
	"ruti-store/module/feature/user/service"
	assistant "ruti-store/utils/assitant"
	"ruti-store/utils/token"
)

var (
	repo   domain.UserRepositoryInterface
	serv   domain.UserServiceInterface
	hand   domain.UserHandlerInterface
	openAi assistant.AssistantServiceInterface
)

func InitializeUser(db *gorm.DB) {
	openAi = assistant.NewAssistantService()
	repo = repository.NewUserRepository(db, openAi)
	serv = service.NewUserService(repo)
	hand = handler.NewUserHandler(serv)
}

func SetupRoutesUser(app *fiber.App, jwt token.JWTInterface, userService domain.UserServiceInterface) {
	api := app.Group("/api/v1/user")
	api.Get("/:id", middleware.AuthMiddleware(jwt, userService), hand.GetUserByID)
	api.Post("/get-profile", middleware.AuthMiddleware(jwt, userService), hand.GetUserProfile)
	api.Post("/edit-profile", middleware.AuthMiddleware(jwt, userService), hand.EditProfile)
	api.Get("/", middleware.AuthMiddleware(jwt, userService), hand.GetAllUser)
	api.Post("/chat-bot", hand.ChatBot)
	api.Delete("/delete/:id", middleware.AuthMiddleware(jwt, userService), hand.DeleteUser)
}
