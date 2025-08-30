package stategraph

import (
	"errors"

	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/memory"
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

func (g *StateGraph) Compile(checkPointer *memory.Memory) (Graph, error) {
	if len(g.nodes) == 0 {
		return Graph{}, errors.New("no nodes in graph")
	}
	// Logic to compile the graph state
	return Graph{
		checkPointer: checkPointer,
		nodes:        g.nodes,
		edges:        g.edges,
		currentNode:  g.nodes[0],
	}, nil
}

const (
	errEdgeReached        = "cannot invoke graph: reached end"
	errMultipleEdgesFound = "multiple edges found, cannot determine next node"
	errBrokenGraph        = "cannot invoke graph: broken edge"
)
