package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
	"ruti-store/module/feature/order/domain"
	"ruti-store/utils/export"
	"ruti-store/utils/response"
	"ruti-store/utils/validator"
	"strconv"
	"time"
)

type OrderHandler struct {
	service domain.OrderServiceInterface
}

func NewOrderHandler(service domain.OrderServiceInterface) domain.OrderHandlerInterface {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}

	searchQuery := c.Query("search")
	filterQuery := c.Query("filter")

	var result []*entities.OrderModels
	var totalItems int64

	switch {
	case searchQuery != "" && filterQuery != "":
		result, totalItems, err = h.service.SearchFilterAndPaginateOrder(currentPage, pageSize, searchQuery, filterQuery)
	case searchQuery != "":
		result, totalItems, err = h.service.SearchAndPaginateOrder(currentPage, pageSize, searchQuery)
	case filterQuery != "":
		result, totalItems, err = h.service.FilterAndPaginateOrder(currentPage, pageSize, filterQuery)
	default:
		result, totalItems, err = h.service.GetAllOrders(currentPage, pageSize)
	}

	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	totalPages, nextPage, prevPage, err := h.service.GetOrdersPage(currentPage, pageSize, int(totalItems))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayOrderSummary(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *OrderHandler) GetAllPayment(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}

	searchQuery := c.Query("search")
	filterQuery := c.Query("filter")

	var result []*entities.OrderModels
	var totalItems int64

	switch {
	case searchQuery != "" && filterQuery != "":
		result, totalItems, err = h.service.SearchFilterAndPaginatePayment(currentPage, pageSize, searchQuery, filterQuery)
	case searchQuery != "":
		result, totalItems, err = h.service.SearchAndPaginateOrder(currentPage, pageSize, searchQuery)
	case filterQuery != "":
		result, totalItems, err = h.service.FilterAndPaginatePayment(currentPage, pageSize, filterQuery)
	default:
		result, totalItems, err = h.service.GetAllOrders(currentPage, pageSize)
	}

	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	totalPages, nextPage, prevPage, err := h.service.GetOrdersPage(currentPage, pageSize, int(totalItems))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayPaymentSummary(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	req := new(domain.CreateOrderRequest)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.service.CreateOrder(currentUser.ID, req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Order created successfully", result)
}

func (h *OrderHandler) Callback(c *fiber.Ctx) error {
	var notificationPayload map[string]interface{}

	if err := json.Unmarshal(c.Body(), &notificationPayload); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to parse notification payload: "+err.Error())
	}

	err := h.service.CallBack(notificationPayload)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildWithoutResponse(c, fiber.StatusOK, "Callback processed successfully")
}

func (h *OrderHandler) CreateCart(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	req := new(domain.CreateCartRequest)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.service.CreateCart(currentUser.ID, req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Cart created successfully", domain.CreateCartFormatter(result))
}

func (h *OrderHandler) DeleteCart(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	id := c.Params("id")
	cartID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	err = h.service.DeleteCartItems(cartID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildWithoutResponse(c, fiber.StatusOK, "Cart deleted successfully")
}

func (h *OrderHandler) GetCartUser(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	result, err := h.service.GetCartUser(currentUser.ID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved get cart", domain.ResponseArrayCart(result))
}

func (h *OrderHandler) CreateOrderCart(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	req := new(domain.CreateOrderCartRequest)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.service.CreateOrderCart(currentUser.ID, req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Order created successfully", result)
}

func (h *OrderHandler) AcceptOrder(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	orderID := c.Params("id")
	if orderID == "" {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	err := h.service.AcceptOrder(orderID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildWithoutResponse(c, fiber.StatusOK, "Accept order successfully")
}

func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	req := new(domain.UpdateOrderStatus)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	err := h.service.UpdateOrderStatus(req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildWithoutResponse(c, fiber.StatusOK, "Update order successfully")
}

func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")
	if orderID == "" {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, err := h.service.GetOrderByID(orderID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Update order successfully", domain.FormatOrderDetail(result))
}

func (h *OrderHandler) GetOrderUser(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}

	filterQuery := c.Query("filter")
	var result []*entities.OrderModels
	var totalItems int64

	if filterQuery != "" {
		result, totalItems, err = h.service.GetAllOrdersWithFilter(currentUser.ID, filterQuery, currentPage, pageSize)
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
		}
	} else {
		result, totalItems, err = h.service.GetAllOrdersByUserID(currentUser.ID, currentPage, pageSize)
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
		}
	}

	totalPages, nextPage, prevPage, err := h.service.GetOrdersPage(currentPage, pageSize, int(totalItems))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayOrderUser(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *OrderHandler) GetCartByID(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}
	id := c.Params("id")
	cartID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, err := h.service.GetCartById(cartID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Success get cart by id", domain.CartFormatter(result))
}

func (h *OrderHandler) GetReportOrder(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	startDateParam := c.Query("start_date")
	endDateParam := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateParam)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid start date format. Use YYYY-MM-DD.")
	}

	endDate, err := time.Parse("2006-01-02", endDateParam)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid end date format. Use YYYY-MM-DD.")
	}

	result, err := h.service.GetReportOrder(startDate, endDate)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	headers := []string{
		"ID Pesanan", "Nama", "Email",
		"Catatan", "Total Kuantitas", "Total Harga",
		"Total Diskon", "Total Bayar", "Status Pesanan",
		"Status Pembayaran", "Dibuat Pada",
	}

	title := fmt.Sprintf("Laporan Pesanan Bulan %s", startDate.Format("January 2006"))

	var data [][]interface{}
	for _, order := range result {
		data = append(data, []interface{}{
			order.IdOrder,
			order.User.Name,
			order.User.Email,
			order.Note,
			order.GrandTotalQuantity,
			order.GrandTotalPrice,
			order.GrandTotalDiscount,
			order.TotalAmountPaid,
			order.OrderStatus,
			order.PaymentStatus,
			order.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	fileName := fmt.Sprintf("Laporan Penjualan_%s.xlsx", startDate.Format("January 2006"))
	err = export.ExportXlsx(c, data, headers, title, fileName)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Error exporting to Xlx: "+err.Error())
	}

	return nil
}
