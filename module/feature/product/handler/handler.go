package handler

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"ruti-store/module/entities"
	"ruti-store/module/feature/product/domain"
	"ruti-store/utils/response"
	"ruti-store/utils/upload"
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
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved product by ID", result)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}
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

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	req := new(domain.UpdateProductRequest)
	id := c.Params("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	err = h.service.UpdateProduct(productID, req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildWithoutResponse(c, fiber.StatusCreated, "Success update product")
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}
	id := c.Params("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	err = h.service.DeleteProduct(productID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildWithoutResponse(c, fiber.StatusOK, "Success delete product")
}

func (h *ProductHandler) GetAllProductsReview(c *fiber.Ctx) error {
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

	result, totalItems, err := h.service.GetProductReviews(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetProductsPage(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayProductReviews(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *ProductHandler) AddPhotoProduct(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}
	req := new(domain.AddPhotoProductRequest)
	file, err := c.FormFile("photo")
	var uploadedURL string
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Error opening file: "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)

		uploadedURL, err = upload.ImageUploadHelper(fileToUpload)
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Error uploading file: "+err.Error())
		}
	}

	req.Photo = uploadedURL

	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.service.AddPhotoProducts(req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusCreated, "Success add photo product", domain.ResponseAddPhotoProduct(result))
}

func (h *ProductHandler) UpdatePhotoProduct(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}
	id := c.Params("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}
	file, err := c.FormFile("photo")
	var uploadedURL string
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Error opening file: "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)

		uploadedURL, err = upload.ImageUploadHelper(fileToUpload)
		if err != nil {
			return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Error uploading file: "+err.Error())
		}
	}

	err = h.service.UpdatePhotoProduct(productID, uploadedURL)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildWithoutResponse(c, fiber.StatusCreated, "Success update photo product")
}
