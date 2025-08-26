GolangGraph is a golang implementation of LangGraph ver 0.3 orchestration framework for building, managing, and deploying long-running, stateful agents.

# Easy to start

```
func main() {
	graphBuilder := stategraph.New()
	llm := gemini.New(gemini.Gemini2_5_Flash, "key")
	graphBuilder.AddNode("chatbot", llm, chatbot)
	graphBuilder.AddEdge(stategraph.EdgeStart, "chatbot")
	graphBuilder.AddEdge("chatbot", stategraph.EdgeEnd)
	memory := inmemorysaver.New()
	graph := graphBuilder.Compile(&memory)
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
		fmt.Println(graph.Invoke(agents.Config{}, agents.Messages{{Role: agents.RoleUser, Content: userInput}}, nil))
	}
}
```

