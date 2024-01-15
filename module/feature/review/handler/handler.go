package handler

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"ruti-store/module/entities"
	"ruti-store/module/feature/review/domain"
	"ruti-store/utils/response"
	"ruti-store/utils/upload"
	"ruti-store/utils/validator"
	"strconv"
)

type ReviewHandler struct {
	service domain.ReviewServiceInterface
}

func NewReviewHandler(service domain.ReviewServiceInterface) domain.ReviewHandlerInterface {
	return &ReviewHandler{
		service: service,
	}
}

func (h *ReviewHandler) GetReviewByID(c *fiber.Ctx) error {
	id := c.Params("id")
	reviewID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, err := h.service.GetReviewById(reviewID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved review by ID", domain.ReviewFormatter(result))
}

func (h *ReviewHandler) GetAllReviewProduct(c *fiber.Ctx) error {
	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}
	id := c.Params("id")
	productID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, totalItems, err := h.service.GetReviewsByProductID(productID, currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetReviewsProductPage(productID, currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayReviews(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *ReviewHandler) CreateReview(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	req := new(domain.CreateReviewRequest)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Failed to parse request body")
	}

	if err := validator.ValidateStruct(req); err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.service.CreateReview(currentUser.ID, req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusCreated, "Success create review", domain.ReviewFormatter(result))
}

func (h *ReviewHandler) CreateReviewPhoto(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "customer" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only customer users can access this resource.")
	}

	req := new(domain.CreatePhotoReviewRequest)
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

	result, err := h.service.CreateReviewImages(req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusCreated, "Success create review", domain.FormatCreateReviewPhotos(result))
}
