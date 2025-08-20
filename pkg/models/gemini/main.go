package gemini

import (
	"context"
	"golanggraph/pkg/models"
	"log"

	"google.golang.org/genai"
)

const (
	Gemini2_5_Flash = "gemini-2.5-flash"
)

type Gemini struct {
	Name string
	Key  string
}

func (g *Gemini) Infer(input string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		g.Name,
		genai.Text(input),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	return result.Text()
}

func New(name, key string) models.Model {
	return &Gemini{
		Name: name,
		Key:  key,
	}
}
