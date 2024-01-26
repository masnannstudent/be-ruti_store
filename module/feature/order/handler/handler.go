package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/entities"
	"ruti-store/module/feature/order/domain"
	"ruti-store/utils/response"
	"ruti-store/utils/validator"
	"strconv"
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

	result, totalItems, err := h.service.GetAllOrders(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetOrdersPage(currentPage, pageSize)
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

	result, totalItems, err := h.service.GetAllOrders(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetOrdersPage(currentPage, pageSize)
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

	result, err := h.service.GetAllOrdersByUserID(currentUser.ID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}
	return response.SuccessBuildResponse(c, fiber.StatusOK, "Success get all order user", domain.FormatterGetAllOrderUser(result))
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
