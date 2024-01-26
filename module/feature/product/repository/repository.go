package repository

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/product/domain"
	assistant "ruti-store/utils/assitant"
	"strings"
)

type ProductRepository struct {
	db     *gorm.DB
	openAi assistant.AssistantServiceInterface
}

func NewProductRepository(db *gorm.DB, openAi assistant.AssistantServiceInterface) domain.ProductRepositoryInterface {
	return &ProductRepository{
		db:     db,
		openAi: openAi,
	}
}

func (r *ProductRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.ProductModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *ProductRepository) GetPaginatedProducts(page, pageSize int) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Order("created_at DESC").
		Offset(offset).Limit(pageSize).Preload("Photos").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(productID uint64) (*entities.ProductModels, error) {
	var product *entities.ProductModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", productID).Preload("Photos").First(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) CreateProduct(newData *entities.ProductModels, categoryIDs []uint64) (*entities.ProductModels, error) {

	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}

	if len(categoryIDs) > 0 {
		categories := make([]entities.CategoryModels, len(categoryIDs))
		for i, categoryID := range categoryIDs {
			categories[i] = entities.CategoryModels{ID: categoryID}
		}

		if err := r.db.Model(newData).Association("Categories").Append(categories); err != nil {
			return nil, err
		}
	}

	return newData, nil
}

func (r *ProductRepository) UpdateProduct(productID uint64, newData *entities.ProductModels, categoryIDs []uint64) error {
	var existingProduct *entities.ProductModels
	if err := r.db.Where("id = ?", productID).First(&existingProduct).Error; err != nil {
		return err
	}

	if err := r.db.Model(&existingProduct).Updates(newData).Error; err != nil {
		return err
	}

	if len(existingProduct.Categories) > 0 {
		if err := r.db.Model(existingProduct).Association("Categories").Delete(existingProduct.Categories); err != nil {
			return err
		}
	}

	if len(categoryIDs) > 0 {
		categories := make([]entities.CategoryModels, len(categoryIDs))
		for i, categoryID := range categoryIDs {
			categories[i] = entities.CategoryModels{ID: categoryID}
		}

		if err := r.db.Model(existingProduct).Association("Categories").Replace(categories); err != nil {
			return err
		}
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(productID uint64) error {
	var existingProduct *entities.ProductModels
	if err := r.db.Where("id = ?", productID).Preload("Categories").First(&existingProduct).Error; err != nil {
		return err
	}

	if len(existingProduct.Categories) > 0 {
		if err := r.db.Model(existingProduct).Association("Categories").Delete(&existingProduct.Categories); err != nil {
			return err
		}
	}

	if err := r.db.Delete(existingProduct).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) UpdateTotalReview(productID uint64) error {
	var products *entities.ProductModels
	err := r.db.Model(&products).Where("id = ?", productID).UpdateColumn("total_reviews", gorm.Expr("total_reviews + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) UpdateProductRating(productID uint64, newRating float64) error {
	if err := r.db.Model(&entities.ProductModels{}).Where("id = ?", productID).Update("rating", newRating).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetProductReviews(page, perPage int) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels
	offset := (page - 1) * perPage
	err := r.db.Where("deleted_at IS NULL").Offset(offset).Limit(perPage).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) AddPhotoProduct(newData *entities.ProductPhotoModels) (*entities.ProductPhotoModels, error) {
	if err := r.db.Create(newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}

func (r *ProductRepository) UpdateProductPhoto(productID uint64, newPhotoURL string) error {
	if err := r.db.Where("product_id = ?", productID).Delete(&entities.ProductPhotoModels{}).Error; err != nil {
		return err
	}

	newPhoto := &entities.ProductPhotoModels{
		ProductID: productID,
		URL:       newPhotoURL,
	}

	if err := r.db.Create(newPhoto).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) ReduceStockWhenPurchasing(productID, quantity uint64) error {
	var products entities.ProductModels
	if err := r.db.Model(&products).Where("id = ?", productID).Update("stock", gorm.Expr("stock - ?", quantity)).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) IncreaseStock(productID, quantity uint64) error {
	var products entities.ProductModels
	if err := r.db.Model(&products).Where("id = ?", productID).Update("stock", gorm.Expr("stock + ?", quantity)).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetAllOrders() ([]*entities.OrderModels, error) {
	var orders []*entities.OrderModels

	if err := r.db.
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Limit(10).
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *ProductRepository) GenerateRecommendationProduct() ([]string, error) {
	ctx := context.Background()

	orders, err := r.GetAllOrders()
	if err != nil {
		return nil, err
	}

	var recommendedProducts []string
	chat := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Anda adalah seorang analis untuk pembelian pengguna.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Berdasarkan pembelian sebelumnya oleh pengguna dan tren terkini di Indonesia, berikan 3 produk yang relevan. Hanya nama produk, tanpa deskripsi atau yang lainnya.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Contoh jawaban\n1. Tas\n2. Baju\n3. Celana",
		},
	}

	orderContent := "Daftar produk yang dibeli oleh pengguna: \n"
	for _, order := range orders {
		for _, product := range order.OrderDetails {
			orderContent += fmt.Sprintf("- %s\n", product.Product.Name)
		}
	}

	chat = append(chat, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: orderContent,
	})

	resp, err := r.openAi.GetAnswerFromAi(chat, ctx)
	if err != nil {
		return nil, err
	}
	for _, choice := range resp.Choices {
		if choice.Message.Role == "assistant" {
			lines := strings.Split(choice.Message.Content, "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) != "" {
					product := strings.SplitN(line, ". ", 2)
					if len(product) > 1 {
						recommendedProducts = append(recommendedProducts, strings.TrimSpace(product[1]))
					}
				}
			}
		}
	}

	return recommendedProducts, nil
}

func (r *ProductRepository) FindAllProductRecommendation(productsFromAI []string) ([]*entities.ProductModels, error) {
	var matchingProducts []*entities.ProductModels

	if len(productsFromAI) == 0 {
		return nil, nil
	}

	var relevantConditions []string
	var relevantValues []interface{}
	for _, productDesc := range productsFromAI {
		relevantConditions = append(relevantConditions, "name LIKE ?")
		relevantValues = append(relevantValues, "%"+productDesc+"%")
	}

	relevantQuery := r.db.Table("product").
		Where(strings.Join(relevantConditions, " OR "), relevantValues...)

	fullQuery := r.db.
		Raw("(?)", relevantQuery).
		Preload("Photos").
		Preload("Categories").
		Limit(3).
		Find(&matchingProducts)

	if fullQuery.Error != nil {
		return nil, fullQuery.Error
	}

	return matchingProducts, nil
}

func (r *ProductRepository) SearchAndPaginateProducts(name string, page, pageSize int) ([]*entities.ProductModels, int64, error) {
	var products []*entities.ProductModels
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Where("name LIKE ?", "%"+name+"%").
		Model(&entities.ProductModels{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Where("name LIKE ?", "%"+name+"%").
		Order("created_at DESC").
		Offset(offset).Limit(pageSize).Preload("Photos").Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, totalItems, nil
}
