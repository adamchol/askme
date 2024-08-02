package api

import (
	"context"

	"github.com/adamchol/askme/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sashabaranov/go-openai"
)

func GetOpenAICompletionStream(cfg openai.ClientConfig, model string, prompt string) tea.Cmd {
	return func() tea.Msg {
		client := openai.NewClientWithConfig(cfg)

		req := openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Stream: true,
		}

		stream, err := client.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			return models.ErrMsg(err)
		}

		return models.InitStreamMsg(stream)
	}
}
