package category

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/category/domain"
	"ruti-store/module/feature/category/handler"
	"ruti-store/module/feature/category/repository"
	"ruti-store/module/feature/category/service"
	"ruti-store/module/feature/middleware"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/token"
)

var (
	repo domain.CategoryRepositoryInterface
	serv domain.CategoryServiceInterface
	hand domain.CategoryHandlerInterface
)

func InitializeCategory(db *gorm.DB) {
	repo = repository.NewCategoryRepository(db)
	serv = service.NewCategoryService(repo)
	hand = handler.NewCategoryHandler(serv)
}

func SetupRoutesCategory(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/category")
	api.Get("/list", hand.GetAllCategories)
	api.Get("/details/:id", hand.GetCategoryByID)
	api.Post("/create", middleware.AuthMiddleware(jwt, userService), hand.CreateCategory)
	api.Put("/update/:id", middleware.AuthMiddleware(jwt, userService), hand.UpdateCategory)
	api.Delete("/delete/:id", middleware.AuthMiddleware(jwt, userService), hand.DeleteCategory)
	api.Get("/product/list/:id", hand.GetAllProductByCategoryID)
}
