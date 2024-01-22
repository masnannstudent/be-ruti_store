package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
	address "ruti-store/module/feature/address/domain"
	addressRepository "ruti-store/module/feature/address/repository"
	addressService "ruti-store/module/feature/address/service"
	"ruti-store/module/feature/middleware"
	notification "ruti-store/module/feature/notification/domain"
	notificationRepository "ruti-store/module/feature/notification/repository"
	notificationService "ruti-store/module/feature/notification/service"
	"ruti-store/module/feature/order/domain"
	"ruti-store/module/feature/order/handler"
	"ruti-store/module/feature/order/repository"
	"ruti-store/module/feature/order/service"
	product "ruti-store/module/feature/product/domain"
	productRepository "ruti-store/module/feature/product/repository"
	productService "ruti-store/module/feature/product/service"
	user "ruti-store/module/feature/user/domain"
	userRepository "ruti-store/module/feature/user/repository"
	userService "ruti-store/module/feature/user/service"
	generator2 "ruti-store/utils/generator"
	"ruti-store/utils/shipping"
	"ruti-store/utils/token"
)

var (
	orderRepo        domain.OrderRepositoryInterface
	orderServ        domain.OrderServiceInterface
	orderHand        domain.OrderHandlerInterface
	productRepo      product.ProductRepositoryInterface
	productServ      product.ProductServiceInterface
	uuidGenerator    generator2.GeneratorInterface
	addressRepo      address.AddressRepositoryInterface
	addressServ      address.AddressServiceInterface
	userRepo         user.UserRepositoryInterface
	userServ         user.UserServiceInterface
	ship             shipping.ShippingServiceInterface
	notificationRepo notification.NotificationRepositoryInterface
	notificationServ notification.NotificationServiceInterface
)

func InitializeOrder(db *gorm.DB, snapClient snap.Client, coreClient coreapi.Client) {
	productRepo = productRepository.NewProductRepository(db)
	productServ = productService.NewProductService(productRepo)
	uuidGenerator = generator2.NewGeneratorUUID(db)
	ship = shipping.NewShippingService()
	addressRepo = addressRepository.NewAddressRepository(db, ship)
	addressServ = addressService.NewAddressService(addressRepo)
	userRepo = userRepository.NewUserRepository(db)
	userServ = userService.NewUserService(userRepo)
	notificationRepo = notificationRepository.NewNotificationRepository(db)
	notificationServ = notificationService.NewNotificationService(notificationRepo)

	orderRepo = repository.NewOrderRepository(db, snapClient, coreClient)
	orderServ = service.NewOrderService(orderRepo, uuidGenerator, productServ, addressServ, userServ, notificationServ)
	orderHand = handler.NewOrderHandler(orderServ)
}

func SetupOrderRoutes(app *fiber.App, jwt token.JWTInterface, userService user.UserServiceInterface) {
	api := app.Group("/api/v1/order")
	api.Get("/payment/list", middleware.AuthMiddleware(jwt, userService), orderHand.GetAllPayment)
	api.Get("/list", middleware.AuthMiddleware(jwt, userService), orderHand.GetAllOrders)
	api.Post("/create", middleware.AuthMiddleware(jwt, userService), orderHand.CreateOrder)
	api.Post("/callback", orderHand.Callback)
	api.Post("/cart/create", middleware.AuthMiddleware(jwt, userService), orderHand.CreateCart)
	api.Delete("/cart/delete/:id", middleware.AuthMiddleware(jwt, userService), orderHand.DeleteCart)
	api.Get("/cart/list", middleware.AuthMiddleware(jwt, userService), orderHand.GetCartUser)
	api.Post("/create/cart", middleware.AuthMiddleware(jwt, userService), orderHand.CreateOrderCart)
	api.Post("/accept/:id", middleware.AuthMiddleware(jwt, userService), orderHand.AcceptOrder)
	api.Put("/update-status", middleware.AuthMiddleware(jwt, userService), orderHand.UpdateOrderStatus)
	api.Get("details/:id", middleware.AuthMiddleware(jwt, userService), orderHand.GetOrderByID)
}
