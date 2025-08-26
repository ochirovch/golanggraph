package stategraph

import (
	"errors"
	"maps"

	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/memory"
	"github.com/ochirovch/golanggraph/pkg/tools"
)

const (
	EdgeStart = "start"
	EdgeEnd   = "end"
)

type Node struct {
	Name     string
	LLM      agents.Invoker
	Function NodeFunc
}

type Edge struct {
	From string
	To   string
}

type NodeFunc func(llm agents.Invoker, messages agents.Messages) (
	retMessages agents.Messages, data map[string]any)

type StateGraph struct {
	nodes []Node
	edges map[string][]string
}

func New() *StateGraph {
	return &StateGraph{
		nodes: make([]Node, 0),
		edges: make(map[string][]string),
	}
}

func (g *StateGraph) AddNode(name string, llm agents.Invoker, fn NodeFunc) {
	g.nodes = append(g.nodes, Node{
		Name:     name,
		LLM:      llm,
		Function: fn,
	})
}

func (g *StateGraph) AddEdge(from, to string) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *StateGraph) Compile(checkPointer *memory.Memory) Graph {
	// Logic to compile the graph state
	return Graph{
		checkPointer: checkPointer,
		nodes:        g.nodes,
		edges:        g.edges,
		currentNode:  string(EdgeStart)}
}

const (
	errEdgeReached        = "cannot invoke graph: reached end"
	errMultipleEdgesFound = "multiple edges found, cannot determine next node"
	errBrokenGraph        = "cannot invoke graph: broken edge"
)

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
