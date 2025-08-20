package react

import (
	"golanggraph/pkg/agents"
	"golanggraph/pkg/memory"
	"golanggraph/pkg/models"
	"golanggraph/pkg/tools"
)

type ReactAgent struct {
	Model  models.Model
	Type   agents.AgentType
	Prompt string // System prompt
	Memory memory.Memory
}

func (r *ReactAgent) Invoke(config agents.Config,
	messages []agents.Message,
	tools []tools.Tool,
) string {
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
	response := r.Model.Infer(prompt)
	// Store the response in memory
	responseMessage := agents.Message{
		Role:    agents.RoleSystem,
		Content: response,
	}
	r.Memory.Store(config.ThreadID, agents.Messages{responseMessage})
	return response
}

func NewAgent(prompt string, model models.Model, memory memory.Memory) agents.Agent {
	return &ReactAgent{
		Model:  model,
		Type:   agents.AgentTypeReact,
		Prompt: prompt,
		Memory: memory,
	}
}
