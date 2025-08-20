package inmemorysaver

import (
	"fmt"
	"golanggraph/pkg/agents"
	"golanggraph/pkg/memory"
)

type InMemorySaver struct {
	store map[string]agents.Messages
}

func New() memory.Memory {
	return &InMemorySaver{
		store: make(map[string]agents.Messages),
	}
}

func (s *InMemorySaver) Store(key string, value []agents.Message) {
	s.store[key] = value
}

func (s *InMemorySaver) Retrieve(key string) ([]agents.Message, error) {
	if value, exists := s.store[key]; exists {
		return value, nil
	}
	return nil, fmt.Errorf("no messages found for key: %s", key)
}
