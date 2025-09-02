GolangGraph is a golang implementation of LangGraph ver 0.3 orchestration framework for building, managing, and deploying long-running, stateful agents.

Goal: implement all examples from [origin](https://langchain-ai.github.io/langgraph/concepts/why-langgraph/)

- [x] check the correctness of the graph when compiling
- [ ] implementing the conditional edges
- [ ] working with tools
- [ ] add human-in-the-loop functionality
- [ ] drawing the graph


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

