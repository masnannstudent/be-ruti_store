package repository

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"ruti-store/module/entities"
	"ruti-store/module/feature/user/domain"
	assistant "ruti-store/utils/assitant"
)

type UserRepository struct {
	db     *gorm.DB
	openAi assistant.AssistantServiceInterface
}

func NewUserRepository(db *gorm.DB, openAi assistant.AssistantServiceInterface) domain.UserRepositoryInterface {
	return &UserRepository{
		db:     db,
		openAi: openAi,
	}
}

func (r *UserRepository) GetUserByID(addressID uint64) (*entities.UserModels, error) {
	var users *entities.UserModels

	if err := r.db.Where("id = ? AND deleted_at IS NULL", addressID).First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) EditProfile(userID uint64, req *entities.UserModels) error {
	var user *entities.UserModels
	if err := r.db.Model(&user).Where("id = ?", userID).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetTotalUserItems() (int64, error) {
	var totalItems int64

	if err := r.db.Where("deleted_at IS NULL").
		Model(&entities.UserModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *UserRepository) GetPaginatedUsers(page, pageSize int) ([]*entities.UserModels, error) {
	var users []*entities.UserModels

	offset := (page - 1) * pageSize

	if err := r.db.Where("deleted_at IS NULL").
		Offset(offset).
		Limit(pageSize).
		Find(&users).
		Order("created_at DESC").
		Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) ChatBotAI(req *domain.CreateChatBotRequest) (string, error) {
	chat := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Nama kamu adalah Ruti Bot",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Kamu adalah seorang chatbot yang bertugas untuk memberikan tips dan saran mengenai dunia fashion",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Kamu adalah hanya boleh menjawab pertanyaan mengenai dunia fashion",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Berikan jawaban maksimal 20 kata",
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: req.Message,
		},
	}

	response, err := r.openAi.GetAnswerFromAi(chat, context.Background())
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", errors.New("no response from the chatbot")
}
