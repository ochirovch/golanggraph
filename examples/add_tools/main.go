package main

import (
	"errors"
	"fmt"

	"encoding/json"

	"github.com/ochirovch/golanggraph/internal/godoc"
	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/agents/state"
	"github.com/ochirovch/golanggraph/pkg/agents/stategraph"
	inmemorysaver "github.com/ochirovch/golanggraph/pkg/memory/InMemorySaver"
	"github.com/ochirovch/golanggraph/pkg/models/gemini"
	"github.com/ochirovch/golanggraph/pkg/tools"
)

func chatbot(llm agents.Invoker, messages agents.Messages) (agents.Messages, map[string]any) {
	retMessages := llm.Invoke(agents.Config{}, messages)
	return retMessages, nil
}

// used when you need to add two numbers
// where is a is the first number and b is the second number
// the result will be the sum of a and b
func Add(inputs map[string]any) (map[string]any, error) {
	a, ok1 := inputs["a"].(int)
	b, ok2 := inputs["b"].(int)
	if !ok1 || !ok2 {
		return nil, errors.New("invalid input types")
	}
	result := a + b
	return map[string]any{"result": result}, nil
}

// used when you need to multiply two numbers
// where is a is the first number and b is the second number
// the result will be the product of a and b
func Multiply(inputs map[string]any) (map[string]any, error) {
	a, ok1 := inputs["a"].(int)
	b, ok2 := inputs["b"].(int)
	if !ok1 || !ok2 {
		return nil, errors.New("invalid input types")
	}
	result := a * b
	return map[string]any{"result": result}, nil
}

type BasicToolNode struct {
	tools []tools.Tool
}

func (btn *BasicToolNode) Call(state state.State) (agents.Messages, error) {
	messages := state.Messages
	if len(messages) == 0 {
		return nil, errors.New("no messages found")
	}
	lastMessage := messages[len(messages)-1]
	var output map[string]any
	var err error
	for _, toolCall := range lastMessage.ToolCalls {
		for _, tool := range btn.tools {
			if godoc.FuncName(tool) == toolCall.Name {
				output, err = tool(toolCall.Args)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	outputJSON, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}
	retMessage := agents.Message{
		Role:    agents.RoleTool,
		Content: string(outputJSON),
	}

	return agents.Messages{retMessage}, nil
}

func NewBasicToolNode(tools []tools.Tool) *BasicToolNode {
	return &BasicToolNode{tools: tools}
}

func main() {
	graphBuilder := stategraph.New()
	llm := gemini.New(gemini.Gemini2_5_Flash, "key")
	llm.BindTools([]tools.Tool{Add, Multiply})
	graphBuilder.AddNode("chatbot", llm, chatbot)
	graphBuilder.AddEdge(stategraph.EdgeStart, "chatbot")
	graphBuilder.AddEdge("chatbot", stategraph.EdgeEnd)
	memory := inmemorysaver.New()
	graph, err := graphBuilder.Compile(&memory)
	if err != nil {
		panic(err)
	}
	response, err := graph.Invoke(agents.Config{
		ThreadID: "thread-1",
	}, agents.Messages{
		{Role: agents.RoleUser, Content: "Hello, how can I use tools?"},
	},
		[]tools.Tool{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
