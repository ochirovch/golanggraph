package main

import (
	"fmt"

	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/agents/stategraph"
	inmemorysaver "github.com/ochirovch/golanggraph/pkg/memory/InMemorySaver"
	"github.com/ochirovch/golanggraph/pkg/models/gemini"
	"github.com/ochirovch/golanggraph/pkg/tools"
)

func chatbot(llm agents.Invoker, messages agents.Messages) (agents.Messages, map[string]any) {
	retMessages := llm.Invoke(agents.Config{}, messages, nil)
	return retMessages, nil
}

func tool_node(llm agents.Invoker, messages agents.Messages) (agents.Messages, map[string]any) {
	retMessages := llm.Invoke(agents.Config{}, messages, nil)
	return retMessages, nil
}

func main() {
	graphBuilder := stategraph.New()
	llm := gemini.New(gemini.Gemini2_5_Flash, "key")
	graphBuilder.AddNode("chatbot", llm, chatbot)
	graphBuilder.AddNode("tool", llm, tool_node)
	graphBuilder.AddEdge(stategraph.EdgeStart, "chatbot")
	graphBuilder.AddEdge("chatbot", stategraph.EdgeEnd)
	memory := inmemorysaver.New()
	graph := graphBuilder.Compile(&memory)
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
