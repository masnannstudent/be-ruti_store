package service

import (
	"errors"
	"math"
	"ruti-store/module/entities"
	"ruti-store/module/feature/article/domain"
	"time"
)

type ArticleService struct {
	repo domain.ArticleRepositoryInterface
}

func NewArticleService(repo domain.ArticleRepositoryInterface) domain.ArticleServiceInterface {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) GetAllArticles(page, pageSize int) ([]*entities.ArticleModels, int64, error) {
	result, err := s.repo.GetPaginatedArticles(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *ArticleService) GetArticlesPage(currentPage, pageSize int) (int, int, int, int, error) {
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

func (s *ArticleService) GetArticleByID(articleID uint64) (*entities.ArticleModels, error) {
	result, err := s.repo.GetArticleByID(articleID)
	if err != nil {
		return nil, errors.New("article not found")
	}
	return result, nil
}

func (s *ArticleService) CreateArticle(req *domain.CreateArticleRequest) (*entities.ArticleModels, error) {
	newData := &entities.ArticleModels{
		Title:     req.Title,
		Content:   req.Content,
		Author:    "Ruti Store",
		Photo:     req.Photo,
		CreatedAt: time.Now(),
	}

	createdArticle, err := s.repo.CreateArticle(newData)
	if err != nil {
		return nil, err
	}
	return createdArticle, nil
}

func (s *ArticleService) UpdateArticle(articleID uint64, req *domain.UpdateArticleRequest) error {
	article, err := s.repo.GetArticleByID(articleID)
	if err != nil {
		return errors.New("article not found")
	}

	newData := &entities.ArticleModels{
		Title:     req.Title,
		Content:   req.Content,
		Photo:     req.Photo,
		UpdatedAt: time.Now(),
	}

	err = s.repo.UpdateArticle(article.ID, newData)
	if err != nil {
		return err
	}

	return nil
}

func (s *ArticleService) DeleteArticle(articleID uint64) error {
	article, err := s.repo.GetArticleByID(articleID)
	if err != nil {
		return errors.New("article not found")
	}

	err = s.repo.DeleteArticle(article.ID)
	if err != nil {
		return err
	}

	return nil
}
