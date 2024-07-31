package services

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/adamchol/askme/internal/api"
	"github.com/charmbracelet/glamour"
	"github.com/liushuangls/go-anthropic/v2"
)

type LLMService struct {
	openAIAPI *api.OpenAIAPI
	claudeAPI *api.ClaudeAPI
}

func NewLLMAIService(config *ConfigService) *LLMService {
	return &LLMService{
		openAIAPI: api.NewOpenAIAPI(config.OpenAI.APIKey),
		claudeAPI: api.NewClaudeAPI(config.Claude.APIKey),
	}
}

func (s *LLMService) ShowGPTMessage(prompt string) error {
	stream, err := s.openAIAPI.GetCompletionStream(prompt)
	if err != nil {
		return err
	}
	defer stream.Close()

	whole := ""
	begin := false
	out := ""

	for {
		resp, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		if begin {
			// sleep()
			for range len(strings.Split(out, "\n")) - 1 {
				fmt.Print("\033[1A")
			}
			for range len(strings.Split(out, "\n")) {
				fmt.Print("\033[9999D")
				fmt.Print("\033[K")
			}

		}

		whole += resp.Choices[0].Delta.Content

		out, _ = glamour.Render(whole, "dark")

		fmt.Print(out)

		begin = true

	}
	return nil
}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}

func (s *LLMService) ShowClaudeMessage(prompt string) error {
	err := s.claudeAPI.GetClaudeMessage(prompt, func(data anthropic.MessagesEventContentBlockDeltaData) {
		fmt.Print(*data.Delta.Text)
	})
	if err != nil {
		return err
	}

	return nil
}
