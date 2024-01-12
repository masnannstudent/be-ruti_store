package response

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type PaginationResponse struct {
	CurrentPage int `json:"current_page"`
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	PrevPage    int `json:"prev_page"`
	NextPage    int `json:"next_page"`
}

type PaginationData struct {
	Message    string             `json:"message"`
	Data       interface{}        `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

func SuccessBuildResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := SuccessResponse{
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}

func ErrorBuildResponse(c *fiber.Ctx, statusCode int, message string) error {
	response := ErrorResponse{
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

func SuccessBuildWithoutResponse(c *fiber.Ctx, statusCode int, message string) error {
	response := ErrorResponse{
		Message: message,
	}
	return c.Status(statusCode).JSON(response)
}

func PaginationBuildResponse(c *fiber.Ctx, statusCode int, message string, data interface{}, currentPage, totalItems, totalPages, nextPage, prevPage int) error {
	pagination := PaginationResponse{
		CurrentPage: currentPage,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		PrevPage:    prevPage,
		NextPage:    nextPage,
	}

	paginationData := PaginationData{
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}

	return c.Status(statusCode).JSON(paginationData)
}
