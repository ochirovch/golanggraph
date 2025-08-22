package react

import (
	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/memory"
	"github.com/ochirovch/golanggraph/pkg/tools"
)

type ReactAgent struct {
	Model  agents.Invoker
	Type   agents.AgentType
	Prompt string // System prompt
	Memory memory.Memory
}

func (r *ReactAgent) Invoke(config agents.Config,
	messages agents.Messages,
	tools []tools.Tool,
) agents.Messages {
	prompt := string(agents.RoleSystem) + ": " + r.Prompt
	// Store the messages in memory
	for _, msg := range messages {
		// convert agents.Message to memory.Message
		memoryMessage := agents.Message{
			Role:    msg.Role,
			Content: msg.Content,
		}
		r.Memory.Store(config.ThreadID, agents.Messages{memoryMessage})
		prompt += "\n" + string(msg.Role) + ": " + msg.Content
	}
	response := r.Model.Invoke(agents.Config{
		ThreadID: config.ThreadID},
		agents.Messages{{
			Role:    agents.RoleUser,
			Content: prompt,
		}},
		nil,
	)
	// Store the response in memory
	r.Memory.Store(config.ThreadID, response)
	return response
}

func NewAgent(prompt string, model agents.Invoker, memory memory.Memory) agents.Invoker {
	return &ReactAgent{
		Model:  model,
		Type:   agents.AgentTypeReact,
		Prompt: prompt,
		Memory: memory,
	}
}
