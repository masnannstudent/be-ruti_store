package assistant

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type AssistantServiceInterface interface {
	GetAnswerFromAi(chat []openai.ChatCompletionMessage, ctx context.Context) (openai.ChatCompletionResponse, error)
}
