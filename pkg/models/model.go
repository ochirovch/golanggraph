package models

import "github.com/ochirovch/golanggraph/pkg/agents"

type AgentType string

const (
	GeminiModel AgentType = "gemini"
)

type Model interface {
	Infer(agents.Messages) agents.Messages
}
