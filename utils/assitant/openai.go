package assistant

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"ruti-store/config"
)

type AssistantService struct {
	client *openai.Client
	apiKey string
}

func NewAssistantService() AssistantServiceInterface {
	apiKey := config.InitConfig().OpenAiKey
	client := openai.NewClient(apiKey)

	return &AssistantService{
		client: client,
		apiKey: apiKey,
	}
}

func (s *AssistantService) GetAnswerFromAi(chat []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error) {
	model := openai.GPT3Dot5Turbo
	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: chat,
		},
	)

	return resp, err
}
