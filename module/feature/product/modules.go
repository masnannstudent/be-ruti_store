package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/middleware"
	"ruti-store/module/feature/product/domain"
	"ruti-store/module/feature/product/handler"
	"ruti-store/module/feature/product/repository"
	"ruti-store/module/feature/product/service"
	user "ruti-store/module/feature/user/domain"
	assistant "ruti-store/utils/assitant"
	"ruti-store/utils/token"
)

var (
	repo   domain.ProductRepositoryInterface
	serv   domain.ProductServiceInterface
	hand   domain.ProductHandlerInterface
	openAi assistant.AssistantServiceInterface
)

func InitializeProduct(db *gorm.DB) {
	openAi = assistant.NewAssistantService()
	repo = repository.NewProductRepository(db, openAi)
	serv = service.NewProductService(repo)
	hand = handler.NewProductHandler(serv)
}

func SetupRoutesProduct(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/product")
	api.Get("/list", hand.GetAllProducts)
	api.Get("/details/:id", hand.GetProductByID)
	api.Post("/create", middleware.AuthMiddleware(jwt, userService), hand.CreateProduct)
	api.Put("/update/:id", middleware.AuthMiddleware(jwt, userService), hand.UpdateProduct)
	api.Delete("/delete/:id", middleware.AuthMiddleware(jwt, userService), hand.DeleteProduct)
	api.Get("/reviews", middleware.AuthMiddleware(jwt, userService), hand.GetAllProductsReview)
	api.Post("/photo/create", middleware.AuthMiddleware(jwt, userService), hand.AddPhotoProduct)
	api.Put("/photo/update/:id", middleware.AuthMiddleware(jwt, userService), hand.UpdatePhotoProduct)
	api.Get("/recommendation", hand.GetProductRecommendation)
	api.Get("/recommendation-user", hand.GetAllProductsRecommendation)
}
