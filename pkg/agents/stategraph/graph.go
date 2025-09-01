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
	currentNode  Node
	data         map[string]any
}

func (g *Graph) calculateCurrentNode() error {
	currentNodeEdges := g.edges[g.currentNode.Name]
	if len(currentNodeEdges) == 0 {
		return errors.New(errEdgeReached)
	}
	if len(currentNodeEdges) > 1 {
		return errors.New(errMultipleEdgesFound)
	}
	var nextNode *Node
	for _, node := range g.nodes {
		if node.Name == currentNodeEdges[0] {
			nextNode = &node
			break
		}
	}
	if nextNode == nil {
		return errors.New(errBrokenGraph)
	}
	if nextNode.Name == string(EdgeEnd) {
		return errors.New(errEdgeReached)
	}
	g.currentNode = *nextNode
	return nil
}

func (g *Graph) prepareMessages(config agents.Config, newMessages agents.Messages) (agents.Messages, error) {
	var messagesToCall agents.Messages
	if g.hasMemory() {
		checkPointer := *g.checkPointer
		oldMessages, err := checkPointer.Retrieve(config.ThreadID)
		if err != nil {
			return agents.Messages{}, err
		}
		messagesToCall = append(messagesToCall, oldMessages...)
		return messagesToCall, nil
	}
	return newMessages, nil
}

func (g Graph) hasMemory() bool {
	return g.checkPointer != nil
}

func (g *Graph) Store(threadID string, messages agents.Messages, data map[string]any) {
	if g.hasMemory() {
		checkPointer := *g.checkPointer
		checkPointer.Store(threadID, messages)

		copiedData := make(map[string]any)
		maps.Copy(copiedData, g.data)
		checkPointer.StoreData(threadID, copiedData)
	}
}

func (g *Graph) Invoke(config agents.Config, newMessages agents.Messages, tools []tools.Tool) (agents.Messages, error) {

	for {
		if g.currentNode.Name == EdgeEnd {
			break
		}
		messagesToCall, err := g.prepareMessages(config, newMessages)
		if err != nil {
			return agents.Messages{}, err
		}
		responseMessages, data := g.currentNode.Function(g.currentNode.LLM, messagesToCall)
		g.Store(config.ThreadID, responseMessages, data)
		if err := g.calculateCurrentNode(); err != nil {
			return agents.Messages{}, err
		}
	}

	return agents.Messages{}, nil
}
