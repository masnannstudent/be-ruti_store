package service

import (
	"errors"
	"fmt"
	"math"
	"ruti-store/module/entities"
	address "ruti-store/module/feature/address/domain"
	notification "ruti-store/module/feature/notification/domain"
	"ruti-store/module/feature/order/domain"
	product "ruti-store/module/feature/product/domain"
	users "ruti-store/module/feature/user/domain"
	"ruti-store/utils/generator"
	"time"
)

type OrderService struct {
	repo                domain.OrderRepositoryInterface
	generatorID         generator.GeneratorInterface
	productService      product.ProductServiceInterface
	addressService      address.AddressServiceInterface
	userService         users.UserServiceInterface
	notificationService notification.NotificationServiceInterface
	//cartService    cart.ServiceCartInterface
}

func NewOrderService(
	repo domain.OrderRepositoryInterface,
	generatorID generator.GeneratorInterface,
	productService product.ProductServiceInterface,
	addressService address.AddressServiceInterface,
	userService users.UserServiceInterface,
	notificationService notification.NotificationServiceInterface,
//cartService cart.ServiceCartInterface,

) domain.OrderServiceInterface {
	return &OrderService{
		repo:                repo,
		generatorID:         generatorID,
		productService:      productService,
		addressService:      addressService,
		userService:         userService,
		notificationService: notificationService,
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

	if err := s.productService.ReduceStockWhenPurchasing(request.ProductID, request.Quantity); err != nil {
		return nil, errors.New("gagal mengurangi stok")
	}

	notificationRequest := domain.CreateNotificationPaymentRequest{
		OrderID:       createdOrder.ID,
		UserID:        createdOrder.UserID,
		PaymentStatus: "Menunggu Konfirmasi",
	}
	_, err = s.SendNotificationPayment(notificationRequest)
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

	for _, orderDetail := range orders.OrderDetails {
		if err := s.productService.IncreaseStock(orderDetail.ProductID, orderDetail.Quantity); err != nil {
			return errors.New("failed to increase stock")
		}
	}

	return nil
}

func (s *OrderService) SendNotificationPayment(request domain.CreateNotificationPaymentRequest) (string, error) {
	var notificationMsg string
	var err error

	user, err := s.userService.GetUserByID(request.UserID)
	if err != nil {
		return "", err
	}
	orders, err := s.repo.GetOrderByID(request.OrderID)
	if err != nil {
		return "", err
	}

	switch request.PaymentStatus {
	case "Menunggu Konfirmasi":
		notificationMsg = fmt.Sprintf("Alloo, %s! Pesananmu dengan ID %s udah berhasil dibuat, nih. Ditunggu yupp!!", user.Name, orders.IdOrder)
	case "Konfirmasi":
		notificationMsg = fmt.Sprintf("Thengkyuu, %s! Pembayaran untuk pesananmu dengan ID %s udah kami terima, nih. Semoga harimu menyenangkan!", user.Name, orders.IdOrder)
	case "Gagal":
		notificationMsg = fmt.Sprintf("Maaf, %s. Pembayaran untuk pesanan dengan ID %s gagal, nih. Beritahu kami apabila kamu butuh bantuan yaa!!", user.Name, orders.IdOrder)
	default:
		return "", errors.New("Status pesanan tidak valid")
	}
	req := &notification.CreateNotificationRequest{
		UserID:  user.ID,
		OrderID: orders.IdOrder,
		Title:   "Status Pembayaran",
		Message: notificationMsg,
	}
	_, err = s.notificationService.CreateNotification(req)
	if err != nil {
		return "", errors.New("error send message")
	}

	return notificationMsg, nil
}

func (s *OrderService) SendNotificationOrder(request domain.CreateNotificationOrderRequest) (string, error) {
	var notificationMsg string
	var err error

	user, err := s.userService.GetUserByID(request.UserID)
	if err != nil {
		return "", err
	}
	orders, err := s.repo.GetOrderByID(request.OrderID)
	if err != nil {
		return "", err
	}

	switch request.OrderStatus {
	case "Pengiriman":
		notificationMsg = fmt.Sprintf("Alloo, %s! Pesanan dengan ID %s udah dalam proses pengiriman, nih. Mohon ditunggu yupp!", user.Name, orders.IdOrder)
	case "Selesai":
		notificationMsg = fmt.Sprintf("Yeayy, %s! Pesananmu dengan ID %s udah sampai tujuan, nih. Semoga sukakk yupp!", user.Name, orders.IdOrder)
	case "Menunggu Konfirmasi":
		notificationMsg = fmt.Sprintf("Alloo, %s! Pesananmu dengan ID %s sedang menunggu konfirmasi, nih. Ditunggu yupp!", user.Name, orders.IdOrder)
	case "Proses":
		notificationMsg = fmt.Sprintf("Alloo, %s! Pesananmu dengan ID %s sedang dalam proses, nih. Ditunggu yupp!", user.Name, orders.IdOrder)
	case "Gagal":
		notificationMsg = fmt.Sprintf("Sowwy, %s. Pesananmu dengan ID %s gagal. Coba lagi, yukk!", user.Name, orders.IdOrder)
	default:
		return "", errors.New("Status pengiriman tidak valid")
	}

	req := &notification.CreateNotificationRequest{
		UserID:  user.ID,
		OrderID: orders.IdOrder,
		Title:   "Status Pesanan",
		Message: notificationMsg,
	}
	_, err = s.notificationService.CreateNotification(req)
	if err != nil {
		return "", errors.New("error send message")
	}

	return notificationMsg, nil
}
