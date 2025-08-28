package stategraph

import (
	"errors"
	"maps"

	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/memory"
	"github.com/ochirovch/golanggraph/pkg/tools"
)

type Graph struct {
	checkPointer *memory.Memory
	nodes        []Node
	edges        map[string][]string
	currentNode  string
	data         map[string]any
}

func (g Graph) Invoke(config agents.Config, messages agents.Messages, tools []tools.Tool) (agents.Messages, error) {

	currentNodeEdges := g.edges[g.currentNode]
	if len(currentNodeEdges) == 0 {
		return agents.Messages{}, errors.New(errEdgeReached)
	}
	if len(currentNodeEdges) > 1 {
		return agents.Messages{}, errors.New(errMultipleEdgesFound)
	}
	var nextNode *Node
	for _, node := range g.nodes {
		if node.Name == currentNodeEdges[0] {
			nextNode = &node
			break
		}
	}
	if nextNode == nil {
		return agents.Messages{}, errors.New(errBrokenGraph)
	}
	if nextNode.Name == string(EdgeEnd) {
		return agents.Messages{}, errors.New(errEdgeReached)
	}
	hasMemory := g.checkPointer != nil
	var checkPointer memory.Memory
	var messagesToCall agents.Messages
	if hasMemory {
		checkPointer = *g.checkPointer
		oldMessages, err := checkPointer.Retrieve(config.ThreadID)
		if err != nil {
			return agents.Messages{}, err
		}
		messagesToCall = append(messagesToCall, oldMessages...)
	}
	responseMessages, data := nextNode.Function(nextNode.LLM, append(messagesToCall, messages...))
	maps.Copy(g.data, data)
	if hasMemory {
		checkPointer.Store(config.ThreadID, responseMessages)
	}
	return responseMessages, nil
}
