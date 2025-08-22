package stategraph

import (
	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/memory"
)

type Graph struct {
	checkPointer *memory.Memory
	graph        map[string]Node
	nodeModels   map[string]agents.Invoker
	currentNode  string
}
