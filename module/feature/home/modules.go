package home

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/home/domain"
	"ruti-store/module/feature/home/handler"
	"ruti-store/module/feature/home/repository"
	"ruti-store/module/feature/home/service"
	"ruti-store/module/feature/middleware"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/token"
)

var (
	repo domain.HomeRepositoryInterface
	serv domain.HomeServiceInterface
	hand domain.HomeHandlerInterface
)

func InitializeHome(db *gorm.DB) {
	repo = repository.NewHomeRepository(db)
	serv = service.NewHomeService(repo)
	hand = handler.NewHomeHandler(serv)
}

func SetupRoutesHome(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/home")
	api.Post("/carousel/create", middleware.AuthMiddleware(jwt, userService), hand.CreateCarousel)
	api.Get("/carousel/details/:id", middleware.AuthMiddleware(jwt, userService), hand.GetCarouselByID)
	api.Get("/carousel/list", hand.GetAllCarouselItems)
	api.Put("/carousel/update/:id", middleware.AuthMiddleware(jwt, userService), hand.UpdateCarousel)
	api.Delete("/carousel/delete/:id", middleware.AuthMiddleware(jwt, userService), hand.DeleteCarousel)
	api.Get("/dashboard", middleware.AuthMiddleware(jwt, userService), hand.GetDashboard)
	api.Get("/latest-order", middleware.AuthMiddleware(jwt, userService), hand.GetAllOrders)
}
