package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	address "ruti-store/module/feature/address/domain"
	"ruti-store/module/feature/order/domain"
	product "ruti-store/module/feature/product/domain"
	users "ruti-store/module/feature/user/domain"
	"ruti-store/utils/generator"
	"time"
)

type OrderService struct {
	repo           domain.OrderRepositoryInterface
	generatorID    generator.GeneratorInterface
	productService product.ProductServiceInterface
	addressService address.AddressServiceInterface
	userService    users.UserServiceInterface
	//cartService    cart.ServiceCartInterface
}

func NewOrderService(
	repo domain.OrderRepositoryInterface,
	generatorID generator.GeneratorInterface,
	productService product.ProductServiceInterface,
	addressService address.AddressServiceInterface,
	userService users.UserServiceInterface,
	//cartService cart.ServiceCartInterface,

) domain.OrderServiceInterface {
	return &OrderService{
		repo:           repo,
		generatorID:    generatorID,
		productService: productService,
		addressService: addressService,
		userService:    userService,
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

func (s *OrderService) CreateOrder(userID uint64, request *domain.CreateOrderRequest) (*domain.CreateOrderResponse, error) {
	orderID, err := s.generatorID.GenerateUUID()
	if err != nil {
		return nil, errors.New("failed to generate order ID")
	}

	idOrder, err := s.generatorID.GenerateOrderID()
	if err != nil {
		return nil, errors.New("failed to generate order ID")
	}

	addresses, err := s.addressService.GetAddressByID(request.AddressID)
	if err != nil {
		return nil, errors.New("address not found")
	}

	products, err := s.productService.GetProductByID(request.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if products.Stock < request.Quantity {
		return nil, errors.New("insufficient stock for this order")
	}

	var orderDetails []entities.OrderDetailsModels
	var totalQuantity, totalPrice, totalDiscount uint64

	orderDetail := entities.OrderDetailsModels{
		OrderID:       orderID,
		ProductID:     request.ProductID,
		Quantity:      request.Quantity,
		TotalPrice:    request.Quantity * (products.Price - products.Discount),
		TotalDiscount: products.Discount * request.Quantity,
	}

	totalQuantity += request.Quantity
	totalPrice += orderDetail.TotalPrice
	totalDiscount += orderDetail.TotalDiscount

	orderDetails = append(orderDetails, orderDetail)

	grandTotalPrice := totalPrice
	totalAmountPaid := grandTotalPrice + 2000

	newData := &entities.OrderModels{
		ID:                 orderID,
		IdOrder:            idOrder,
		AddressID:          addresses.ID,
		UserID:             userID,
		Note:               request.Note,
		GrandTotalQuantity: totalQuantity,
		GrandTotalPrice:    grandTotalPrice,
		ShipmentFee:        0,
		AdminFees:          2000,
		TotalAmountPaid:    totalAmountPaid,
		OrderStatus:        "Menunggu Konfirmasi",
		PaymentStatus:      "Menunggu Konfirmasi",
		CreatedAt:          time.Now(),
		OrderDetails:       orderDetails,
	}

	createdOrder, err := s.repo.CreateOrder(newData)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetUserByID(createdOrder.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	snapResult, err := s.repo.CreateSnap(createdOrder.ID, user.Name, user.Email, createdOrder.TotalAmountPaid)
	if err != nil {
		return nil, err
	}

	response := &domain.CreateOrderResponse{
		OrderID:         createdOrder.ID,
		IdOrder:         createdOrder.IdOrder,
		RedirectURL:     snapResult.RedirectURL,
		TotalAmountPaid: createdOrder.TotalAmountPaid,
	}
	return response, nil
}

func (s *OrderService) CallBack(req map[string]interface{}) error {
	orderID, exist := req["order_id"].(string)
	if !exist {
		return errors.New("invalid notification payload")
	}

	status, err := s.repo.CheckTransaction(orderID)
	if err != nil {
		return err
	}

	transaction, err := s.repo.GetOrderByID(orderID)
	if err != nil {
		return errors.New("transaction data not found")
	}

	if status.PaymentStatus == "Konfirmasi" {
		if err := s.ConfirmPayment(transaction.ID); err != nil {
			return err
		}
	} else if status.PaymentStatus == "Gagal" {
		if err := s.CancelPayment(transaction.ID); err != nil {
			return err
		}
	}

	return nil
}

func (s *OrderService) GetOrderByID(orderID string) (*entities.OrderModels, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OrderService) ConfirmPayment(orderID string) error {
	orders, err := s.repo.GetOrderByID(orderID)
	if err != nil {
		return errors.New("pesanan tidak ditemukan")
	}

	orders.OrderStatus = "Proses"
	orders.PaymentStatus = "Konfirmasi"

	if err := s.repo.UpdatePayment(orders.ID, orders.OrderStatus, orders.PaymentStatus); err != nil {
		return err
	}

	return nil
}

func (s *OrderService) CancelPayment(orderID string) error {
	orders, err := s.repo.GetOrderByID(orderID)
	if err != nil {
		return errors.New("pesanan tidak ditemukan")
	}

	orders.OrderStatus = "Gagal"
	orders.PaymentStatus = "Gagal"

	if err := s.repo.UpdatePayment(orderID, orders.OrderStatus, orders.PaymentStatus); err != nil {
		return err
	}

	return nil
}
