package memory

import "github.com/ochirovch/golanggraph/pkg/agents"

type Memory interface {
	Store(key string, value []agents.Message)
	StoreData(key string, data map[string]any)
	Retrieve(key string) ([]agents.Message, error)
}
