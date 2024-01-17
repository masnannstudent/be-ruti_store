package notification

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/notification/domain"
	"ruti-store/module/feature/notification/handler"
	"ruti-store/module/feature/notification/repository"
	"ruti-store/module/feature/notification/service"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/token"
)

var (
	repo domain.NotificationRepositoryInterface
	serv domain.NotificationServiceInterface
	hand domain.NotificationHandlerInterface
)

func InitializeNotification(db *gorm.DB) {
	repo = repository.NewNotificationRepository(db)
	serv = service.NewNotificationService(repo)
	hand = handler.NewNotificationHandler(serv)
}

func SetupRoutesNotification(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	//api := app.Group("/api/v1/notification")
	//api.Get("/list", middleware.AuthMiddleware(jwt, userService), hand.GetNotifications)
}
