package memory

import "github.com/ochirovch/golanggraph/pkg/agents"

type Memory interface {
	Store(key string, value []agents.Message)
	Retrieve(key string) ([]agents.Message, error)
}
