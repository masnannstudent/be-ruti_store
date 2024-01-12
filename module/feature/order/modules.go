package order

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"ruti-store/module/feature/order/domain"
	"ruti-store/module/feature/order/handler"
	"ruti-store/module/feature/order/repository"
	"ruti-store/module/feature/order/service"
	product "ruti-store/module/feature/product/domain"
	productRepository "ruti-store/module/feature/product/repository"
	productService "ruti-store/module/feature/product/service"
	generator2 "ruti-store/utils/generator"
)

var (
	orderRepo     domain.OrderRepositoryInterface
	orderServ     domain.OrderServiceInterface
	orderHand     domain.OrderHandlerInterface
	productRepo   product.ProductRepositoryInterface
	productServ   product.ProductServiceInterface
	uuidGenerator generator2.GeneratorInterface
)

func InitializeOrder(db *gorm.DB) {
	productRepo = productRepository.NewProductRepository(db)
	productServ = productService.NewProductService(productRepo)
	uuidGenerator = generator2.NewGeneratorUUID(db)

	orderRepo = repository.NewOrderRepository(db)
	orderServ = service.NewOrderService(orderRepo, uuidGenerator, productServ)
	orderHand = handler.NewOrderHandler(orderServ)
}

func SetupOrderRoutes(app *fiber.App) {
	api := app.Group("/api/v1/orders")
	api.Get("", orderHand.GetAllOrders)
	api.Get("/payment", orderHand.GetAllPayment)
}
