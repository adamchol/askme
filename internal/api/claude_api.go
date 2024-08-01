package api

import (
	"context"

	"github.com/liushuangls/go-anthropic/v2"
)

type ClaudeAPI struct {
	apiKey string
}

func NewClaudeAPI(apiKey string) *ClaudeAPI {
	return &ClaudeAPI{
		apiKey: apiKey,
	}
}

func (api *ClaudeAPI) GetClaudeMessage(prompt string, callback func(data anthropic.MessagesEventContentBlockDeltaData)) error {
	client := anthropic.NewClient(api.apiKey)
	_, err := client.CreateMessagesStream(context.Background(), anthropic.MessagesStreamRequest{
		MessagesRequest: anthropic.MessagesRequest{
			Model: anthropic.ModelClaude3Dot5Sonnet20240620,
			Messages: []anthropic.Message{
				anthropic.NewUserTextMessage(prompt),
			},
			MaxTokens: 4096,
		},
		OnContentBlockDelta: callback,
	})
	if err != nil {
		return err
	}

	return nil
}
