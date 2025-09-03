package stategraph

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/agents/node"
	"github.com/ochirovch/golanggraph/pkg/agents/state"
	"github.com/ochirovch/golanggraph/pkg/memory"

	"github.com/ochirovch/golanggraph/pkg/tools"
)

type ThreadIDStep struct {
	ThreadID string
	Step     int
}

type Graph struct {
	checkPointer *memory.Memory
	nodes        []node.Node
	edges        map[string][]string
	branches     map[string]conditionalEdge
	states       map[ThreadIDStep]state.State
	currentStep  int
	currentNode  node.Node
}

func (g *Graph) calculateCurrentNode(config agents.Config) error {
	currentNodeEdges := g.edges[g.currentNode.Name]
	if len(currentNodeEdges) == 0 {
		return errors.New(errEdgeReached)
	}
	if len(currentNodeEdges) > 1 {
		return errors.New(errMultipleEdgesFound)
	}
	var nextNode *node.Node
	for _, n := range g.nodes {
		if n.Name == currentNodeEdges[0] {
			nextNode = &n
			break
		}
	}
	if nextNode == nil {
		// If no next node is found, use the conditional edge function
		conditionalNextNode := g.branches[g.currentNode.Name].fn
		path := conditionalNextNode(g.states[ThreadIDStep{
			ThreadID: config.ThreadID,
			Step:     g.currentStep,
		}])
		nextNodeName := g.branches[g.currentNode.Name].pathMap[path]
		for _, n := range g.nodes {
			if n.Name == nextNodeName {
				nextNode = &n
				break
			}
		}
		if nextNode == nil {
			return errors.New(errBrokenGraph)
		}
	}
	if nextNode.Name == string(EdgeEnd) {
		return errors.New(errEdgeReached)
	}
	g.currentNode = *nextNode
	g.currentStep++
	return nil
}

func (g *Graph) prepareMessages(config agents.Config, newMessages agents.Messages) (agents.Messages, error) {
	var messagesToCall agents.Messages
	if g.hasMemory() {
		checkPointer := *g.checkPointer
		oldStates, err := checkPointer.Restore(config.ThreadID)
		if err != nil {
			return agents.Messages{}, err
		}
		for _, state := range oldStates {
			messagesToCall = append(messagesToCall, state.Messages...)
		}
		return messagesToCall, nil
	}
	return newMessages, nil
}

func (g Graph) hasMemory() bool {
	return g.checkPointer != nil
}

func (g *Graph) store(threadID string, messages agents.Messages, data map[string]any) {
	state := state.State{
		ThreadID:    threadID,
		UUID:        uuid.New(),
		Step:        g.currentStep,
		CurrentNode: g.currentNode,
		Messages:    messages,
		Data:        data,
	}
	g.states[ThreadIDStep{
		ThreadID: threadID,
		Step:     g.currentStep,
	}] = state
	if g.hasMemory() {
		checkPointer := *g.checkPointer
		checkPointer.Save(state)
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
		g.store(config.ThreadID, responseMessages, data)
		if err := g.calculateCurrentNode(config); err != nil {
			return agents.Messages{}, err
		}
	}

	return agents.Messages{}, nil
}
