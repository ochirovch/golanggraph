package gemini

import (
	"context"
	"log"

	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/tools"

	"google.golang.org/genai"
)

const (
	Gemini2_5_Flash = "gemini-2.5-flash"
)

type Gemini struct {
	Name string
	Key  string
}

func (g *Gemini) Invoke(config agents.Config, input agents.Messages, tools []tools.Tool) agents.Messages {
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
	return agents.Messages{{Role: agents.RoleAssistant, Content: result.Text()}}
}

func New(name, key string) agents.Invoker {
	return &Gemini{
		Name: name,
		Key:  key,
	}
}
