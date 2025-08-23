package stategraph

import (
	"github.com/ochirovch/golanggraph/pkg/memory"
)

type Graph struct {
	checkPointer *memory.Memory
	nodes        []Node
	edges        map[string][]string
	currentNode  string
}
