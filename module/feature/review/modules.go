package review

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/review/domain"
	"ruti-store/module/feature/review/handler"
	"ruti-store/module/feature/review/repository"
	"ruti-store/module/feature/review/service"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/token"
)

var (
	reviewRepo domain.ReviewRepositoryInterface
	reviewServ domain.ReviewServiceInterface
	reviewHand domain.ReviewHandlerInterface
)

func InitializeReviews(db *gorm.DB) {
	reviewRepo = repository.NewReviewRepository(db)
	reviewServ = service.NewReviewService(reviewRepo)
	reviewHand = handler.NewReviewHandler(reviewServ)
}

func SetupRoutesReviews(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/reviews")
	api.Get("/list", reviewHand.GetAllReviews)
	api.Get("/details/:id", reviewHand.GetReviewByID)
}
