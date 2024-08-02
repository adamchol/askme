package internal

import (
	"errors"
	"io"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sashabaranov/go-openai"
)

type UIModel struct {
	output string
	prompt string
	model  string
	state  state
	error  error
}

type state int

const (
	responseState state = iota
	doneState
	errState
)

type errMsg error

type deltaMsg struct {
	content string
	stream  *openai.ChatCompletionStream
}

type initStreamMsg *openai.ChatCompletionStream

func (m *UIModel) Init() tea.Cmd {
	return m.getOpenAICompletionStream(openai.DefaultConfig(""), m.model, m.prompt)
}

func (m *UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyType:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case errMsg:
		m.error = msg
		m.state = errState
		return m, tea.Quit
	case initStreamMsg:
		m.state = responseState
		return m, m.streamCompletion(msg)
	case deltaMsg:
		if msg.stream == nil {
			m.state = doneState
			return m, tea.Quit
		}

		m.output += msg.content
		return m, m.streamCompletion(msg.stream)
	}
	return m, nil
}

func (m *UIModel) View() string {
	if m.state == errState {
		return "Error occured"
	}
	return m.output
}

func (m *UIModel) streamCompletion(stream *openai.ChatCompletionStream) tea.Cmd {
	return func() tea.Msg {
		resp, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			stream.Close()
			return deltaMsg{}
		}

		if err != nil {
			return errMsg(err)
		}

		return deltaMsg{
			stream:  stream,
			content: resp.Choices[0].Delta.Content,
		}

	}
}
