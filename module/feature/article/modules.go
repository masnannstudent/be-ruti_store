package article

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/article/domain"
	"ruti-store/module/feature/article/handler"
	"ruti-store/module/feature/article/repository"
	"ruti-store/module/feature/article/service"
	"ruti-store/module/feature/middleware"
	user "ruti-store/module/feature/user/domain"
	"ruti-store/utils/token"
)

var (
	repo domain.ArticleRepositoryInterface
	serv domain.ArticleServiceInterface
	hand domain.ArticleHandlerInterface
)

func InitializeArticle(db *gorm.DB) {
	repo = repository.NewArticleRepository(db)
	serv = service.NewArticleService(repo)
	hand = handler.NewArticleHandler(serv)
}

func SetupRoutesArticle(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/article")
	api.Get("/list", hand.GetAllArticles)
	api.Get("/details/:id", hand.GetArticleByID)
	api.Post("/create", middleware.AuthMiddleware(jwt, userService), hand.CreateArticle)
}
