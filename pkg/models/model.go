package models

type AgentType string

const (
	GeminiModel AgentType = "gemini"
)

type Model interface {
	Infer(string) string
}
