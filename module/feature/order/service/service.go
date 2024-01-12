package service

import (
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/order/domain"
	product "ruti-store/module/feature/product/domain"
	"ruti-store/utils/generator"
)

type OrderService struct {
	repo           domain.OrderRepositoryInterface
	generatorID    generator.GeneratorInterface
	productService product.ProductServiceInterface
	//addressService address.ServiceAddressInterface
	//userService    users.ServiceUserInterface
	//cartService    cart.ServiceCartInterface
}

func NewOrderService(
	repo domain.OrderRepositoryInterface,
	generatorID generator.GeneratorInterface,
	productService product.ProductServiceInterface,
	//addressService address.ServiceAddressInterface,
	//userService users.ServiceUserInterface,
	//cartService cart.ServiceCartInterface,

) domain.OrderServiceInterface {
	return &OrderService{
		repo:           repo,
		generatorID:    generatorID,
		productService: productService,
		//addressService: addressService,
		//userService:    userService,
		//cartService:    cartService,
	}
}

func (s *OrderService) GetAllOrders(page, pageSize int) ([]*entities.OrderModels, int64, error) {
	result, err := s.repo.GetPaginatedOrders(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *OrderService) GetOrdersPage(currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return currentPage, totalPages, nextPage, prevPage, nil
}
