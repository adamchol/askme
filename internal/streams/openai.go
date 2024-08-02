package streams

import (
	"errors"
	"io"

	"github.com/adamchol/askme/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sashabaranov/go-openai"
)

func StreamOpenAICompletion(stream *openai.ChatCompletionStream) tea.Cmd {
	return func() tea.Msg {
		resp, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			stream.Close()
			return models.DeltaMsg{}
		}

		if err != nil {
			return models.ErrMsg(err)
		}

		return models.DeltaMsg{
			Stream:  stream,
			Content: resp.Choices[0].Delta.Content,
		}

	}
}
