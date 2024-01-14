package handler

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"ruti-store/module/entities"
	"ruti-store/module/feature/home/domain"
	"ruti-store/utils/response"
	"ruti-store/utils/upload"
	"ruti-store/utils/validator"
	"strconv"
)

type HomeHandler struct {
	service domain.HomeServiceInterface
}

func NewHomeHandler(service domain.HomeServiceInterface) domain.HomeHandlerInterface {
	return &HomeHandler{
		service: service,
	}
}

func (h *HomeHandler) CreateCarousel(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}
	req := new(domain.CreateCarouselRequest)
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

	result, err := h.service.CreateCarousel(req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusCreated, "Success create carousels", domain.CarouselFormatter(result))

}

func (h *HomeHandler) GetCarouselByID(c *fiber.Ctx) error {
	id := c.Params("id")
	carouselID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, err := h.service.GetCarouselById(carouselID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to retrieve carousel: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved carousel by ID", domain.CarouselFormatter(result))
}

func (h *HomeHandler) GetAllCarouselItems(c *fiber.Ctx) error {
	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}

	result, totalItems, err := h.service.GetAllCarouselItems(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetCarouselPage(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayCarousel(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *HomeHandler) UpdateCarousel(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	id := c.Params("id")
	carouselID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	req := new(domain.UpdateCarouselRequest)

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

	err = h.service.UpdateCarousel(carouselID, req)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildWithoutResponse(c, fiber.StatusOK, "Success update carousels")

}

func (h *HomeHandler) DeleteCarousel(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	id := c.Params("id")
	carouselID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	err = h.service.DeleteCarousel(carouselID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to retrieve carousel: "+err.Error())
	}

	return response.SuccessBuildWithoutResponse(c, fiber.StatusOK, "Success delete carousels")
}

func (h *HomeHandler) GetDashboard(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("currentUser").(*entities.UserModels)
	if !ok || currentUser == nil {
		return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Missing or invalid user information.")
	}

	if currentUser.Role != "admin" {
		return response.ErrorBuildResponse(c, fiber.StatusForbidden, "Forbidden: Only admin users can access this resource.")
	}

	totalIncome, totalProduct, totalUser, err := h.service.GetDashboardPage()
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to retrieve dashboard: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved dashboard",
		domain.FormatDashboardResponse(totalIncome, totalProduct, totalUser))

}
