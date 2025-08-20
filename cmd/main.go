package main

import (
	"golanggraph/pkg/agents"
	"golanggraph/pkg/agents/react"
	inmemorysaver "golanggraph/pkg/memory/InMemorySaver"
	"golanggraph/pkg/models/gemini"
	"golanggraph/pkg/tools"
)

func main() {

	// Initialize the in-memory saver
	memory := inmemorysaver.New()
	model := gemini.New(gemini.Gemini2_5_Flash, "key")
	agent := react.NewAgent("system prompt", model, memory)
	config := agents.Config{
		ThreadID: "thread-1",
	}
	response := agent.Invoke(
		config,
		[]agents.Message{
			{Role: agents.RoleUser, Content: "Hello, how are you?"},
		},
		[]tools.Tool{},
	)

	// Print the response
	println(response)
}
