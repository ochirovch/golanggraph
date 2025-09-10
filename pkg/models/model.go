package models

import (
	"github.com/ochirovch/golanggraph/pkg/agents/message"
)

type AgentType string

const (
	GeminiModel AgentType = "gemini"
)

type Model interface {
	Infer(message.Messages) message.Message
}
