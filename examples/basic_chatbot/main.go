package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ochirovch/golanggraph/pkg/agents"
	inmemorysaver "github.com/ochirovch/golanggraph/pkg/memory/InMemorySaver"
	"github.com/ochirovch/golanggraph/pkg/models/gemini"

	"github.com/ochirovch/golanggraph/pkg/agents/invoker"
	"github.com/ochirovch/golanggraph/pkg/agents/message"
	"github.com/ochirovch/golanggraph/pkg/agents/stategraph"
)

func chatbot(llm invoker.Invoker, messages message.Messages) (message.Messages, map[string]any) {
	retMessages := llm.Invoke(agents.Config{}, messages)
	return retMessages, nil
}

func main() {
	graphBuilder := stategraph.New()
	llm := gemini.New(gemini.Gemini2_5_Flash, "key")
	graphBuilder.AddNode("chatbot", llm, chatbot)
	graphBuilder.AddEdge(stategraph.EdgeStart, "chatbot")
	graphBuilder.AddEdge("chatbot", stategraph.EdgeEnd)
	memory := inmemorysaver.New()
	graph, err := graphBuilder.Compile(&memory)
	if err != nil {
		fmt.Println("Error compiling graph:", err)
		return
	}
	for {
		var userInput string
		// ask the user for input
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			userInput = scanner.Text()
		}

		// process user input
		switch strings.ToLower(userInput) {
		case "quit", "exit", "q":
			return
		}
		fmt.Println(graph.Invoke(agents.Config{}, message.Messages{{Role: agents.RoleUser, Content: userInput}}, nil))
	}
}
