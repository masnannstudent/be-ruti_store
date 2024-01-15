package handler

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/feature/product/domain"
	"ruti-store/utils/response"
	"ruti-store/utils/validator"
	"strconv"
)

type ProductHandler struct {
	service domain.ProductServiceInterface
}

func NewProductHandler(service domain.ProductServiceInterface) domain.ProductHandlerInterface {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}

	result, totalItems, err := h.service.GetAllProducts(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetProductsPage(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayProducts(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, err := h.service.GetProductByID(productID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to retrieve product: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved product by ID", result)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	req := new(domain.CreateProductRequest)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.service.CreateProduct(req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusCreated, "Success create product", domain.ResponseDetailProducts(result))
}
