package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/product/domain"
	"ruti-store/module/feature/product/handler"
	"ruti-store/module/feature/product/repository"
	"ruti-store/module/feature/product/service"
)

var (
	repo domain.ProductRepositoryInterface
	serv domain.ProductServiceInterface
	hand domain.ProductHandlerInterface
)

func InitializeProduct(db *gorm.DB) {
	repo = repository.NewProductRepository(db)
	serv = service.NewProductService(repo)
	hand = handler.NewProductHandler(serv)
}

func SetupRoutesProduct(app *fiber.App) {
	api := app.Group("/api/v1/product")
	api.Get("", hand.GetAllProducts)
	api.Get("/:id", hand.GetProductByID)
}
