package internal

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sashabaranov/go-openai"
)

func (m *UIModel) getOpenAICompletionStream(cfg openai.ClientConfig, model string, prompt string) tea.Cmd {
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
			return errMsg(err)
		}

		return initStreamMsg(stream)
	}
}
