package models

import "github.com/sashabaranov/go-openai"

type ErrMsg error

type DeltaMsg struct {
	Content string
	Stream  *openai.ChatCompletionStream
}

type InitStreamMsg *openai.ChatCompletionStream
