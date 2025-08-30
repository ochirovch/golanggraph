package inmemorysaver

import (
	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/memory"
)

type InMemorySaver struct {
	store     map[string]agents.Messages
	storeData map[string]any
}

func New() memory.Memory {
	return &InMemorySaver{
		store: make(map[string]agents.Messages),
	}
}

func (s *InMemorySaver) Store(key string, value []agents.Message) {
	s.store[key] = value
}

func (s *InMemorySaver) StoreData(key string, data map[string]any) {
	s.storeData[key] = data
}

func (s *InMemorySaver) Retrieve(key string) ([]agents.Message, error) {
	if value, exists := s.store[key]; exists {
		return value, nil
	}
	return []agents.Message{}, nil
}
