package gemini

import (
	"context"
	"log"

	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/agents/invoker"
	"github.com/ochirovch/golanggraph/pkg/agents/message"
	"github.com/ochirovch/golanggraph/pkg/agents/tools"

	"google.golang.org/genai"
)

const (
	Gemini2_5_Flash = "gemini-2.5-flash"
)

type Gemini struct {
	Name  string
	Key   string
	Tools []tools.Tool
}

func (g *Gemini) Invoke(config agents.Config, input message.Messages) message.Messages {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	inputText := input.Print()

	result, err := client.Models.GenerateContent(
		ctx,
		g.Name,
		genai.Text(inputText),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return message.Messages{{Role: agents.RoleAssistant, Content: result.Text()}}
}

func (g *Gemini) BindTools(tools []tools.Tool) {
	// Bind tools to the Gemini instance
	g.Tools = tools
}

func (g *Gemini) GetTools() []tools.Tool {
	return g.Tools
}

func New(name, key string) invoker.Invoker {
	return &Gemini{
		Name: name,
		Key:  key,
	}
}
