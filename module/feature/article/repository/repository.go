package repository

import (
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/article/domain"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) domain.ArticleRepositoryInterface {
	return &ArticleRepository{
		db: db,
	}
}

func (r *ArticleRepository) GetPaginatedArticles(page, pageSize int) ([]*entities.ArticleModels, error) {
	var articles []*entities.ArticleModels

	offset := (page - 1) * pageSize

	if err := r.db.
		Where("deleted_at IS NULL").
		Offset(offset).Limit(pageSize).
		Order("created_at DESC").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.ArticleModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *ArticleRepository) GetArticleByID(articleID uint64) (*entities.ArticleModels, error) {
	var article *entities.ArticleModels

	if err := r.db.
		Where("id = ? AND deleted_at IS NULL", articleID).
		First(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (r *ArticleRepository) CreateArticle(article *entities.ArticleModels) (*entities.ArticleModels, error) {
	if err := r.db.Create(article).Error; err != nil {
		return nil, err
	}

	return article, nil
}
