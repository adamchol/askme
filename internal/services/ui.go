package services

import (
	"github.com/adamchol/askme/internal"
	"github.com/adamchol/askme/internal/api"
	. "github.com/adamchol/askme/internal/models"
	"github.com/adamchol/askme/internal/streams"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/sashabaranov/go-openai"
)

type UIModel struct {
	output string
	Input  CompletionInput
	state  State
	error  error

	Config internal.Config
}

type CompletionInput struct {
	Prompt string
	Model  string
}

func (m *UIModel) Init() tea.Cmd {
	return api.GetOpenAICompletionStream(openai.DefaultConfig(m.Config.OpenAI.APIKey), m.Input.Model, m.Input.Prompt)
}

func (m *UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyType:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case ErrMsg:
		m.error = msg
		m.state = ErrState
		return m, tea.Quit
	case InitStreamMsg:
		m.state = ResponseState
		return m, streams.StreamOpenAICompletion(msg)
	case DeltaMsg:
		if msg.Stream == nil {
			m.state = DoneState
			return m, tea.Quit
		}

		m.output += msg.Content
		return m, streams.StreamOpenAICompletion(msg.Stream)
	}
	return m, nil
}

func (m *UIModel) View() string {
	var s string
	if m.state == ErrState {
		s = "Error occured"
		return s
	}

	md, _ := glamour.Render(m.output, "dark")

	s = md

	if m.state == DoneState {
		s += "\n"
	}
	return s
}
