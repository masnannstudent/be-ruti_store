package handler

import (
	"github.com/gofiber/fiber/v2"
	"ruti-store/module/feature/article/domain"
	"ruti-store/utils/response"
	"strconv"
)

type ArticleHandler struct {
	service domain.ArticleServiceInterface
}

func NewArticleHandler(service domain.ArticleServiceInterface) domain.ArticleHandlerInterface {
	return &ArticleHandler{
		service: service,
	}
}

func (h *ArticleHandler) GetAllArticles(c *fiber.Ctx) error {
	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}

	result, totalItems, err := h.service.GetAllArticles(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetArticlesPage(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination",
		domain.ResponseArrayArticles(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}

func (h *ArticleHandler) GetArticleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	articleID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid input format.")
	}

	result, err := h.service.GetArticleByID(articleID)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Internal server error occurred: "+err.Error())
	}

	return response.SuccessBuildResponse(c, fiber.StatusOK, "Successfully retrieved article by ID", domain.ArticleDetailFormatter(result))
}
