package api

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAIAPI struct {
	apiKey string
}

func NewOpenAIAPI(apiKey string) *OpenAIAPI {
	return &OpenAIAPI{
		apiKey: apiKey,
	}
}

func (api *OpenAIAPI) GetCompletionStream(prompt string) (*openai.ChatCompletionStream, error) {

	client := openai.NewClient(api.apiKey)

	req := openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Stream: true,
	}

	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		req,
	)

	if err != nil {
		return nil, err
	}

	return stream, nil
}
