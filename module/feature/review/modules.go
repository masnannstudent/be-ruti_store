package review

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/middleware"
	products "ruti-store/module/feature/product/domain"
	productsRepo "ruti-store/module/feature/product/repository"
	productsService "ruti-store/module/feature/product/service"
	"ruti-store/module/feature/review/domain"
	"ruti-store/module/feature/review/handler"
	"ruti-store/module/feature/review/repository"
	"ruti-store/module/feature/review/service"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/token"
)

var (
	reviewRepo  domain.ReviewRepositoryInterface
	reviewServ  domain.ReviewServiceInterface
	reviewHand  domain.ReviewHandlerInterface
	productServ products.ProductServiceInterface
	productRepo products.ProductRepositoryInterface
)

func InitializeReviews(db *gorm.DB) {
	reviewRepo = repository.NewReviewRepository(db)
	productRepo = productsRepo.NewProductRepository(db)
	productServ = productsService.NewProductService(productRepo)
	reviewServ = service.NewReviewService(reviewRepo, productServ)
	reviewHand = handler.NewReviewHandler(reviewServ)
}

func SetupRoutesReviews(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/reviews")
	api.Get("/details/:id", reviewHand.GetReviewByID)
	api.Get("list/:id", reviewHand.GetAllReviewProduct)
	api.Post("/create", middleware.AuthMiddleware(jwt, userService), reviewHand.CreateReview)
	api.Post("/create/photos", middleware.AuthMiddleware(jwt, userService), reviewHand.CreateReviewPhoto)
}
