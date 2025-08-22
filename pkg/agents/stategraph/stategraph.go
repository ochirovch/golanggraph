package stategraph

import (
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
	Function NodeFunc
}

type Edge struct {
	From string
	To   string
}

type NodeFunc func(llm agents.Invoker, messages agents.Messages) agents.Messages

type StateGraph struct {
	nodes      map[string]Node
	edges      map[string][]string
	nodeModels map[string]agents.Invoker
}

func New() *StateGraph {
	return &StateGraph{
		nodes:      make(map[string]Node),
		nodeModels: make(map[string]agents.Invoker),
	}
}

func (g *StateGraph) AddNode(node string, handler agents.Invoker, fn NodeFunc) {
	g.nodes[node] = Node{
		Name:     node,
		Function: fn,
	}
	g.nodeModels[node] = handler
}

func (g *StateGraph) AddEdge(from, to string) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *StateGraph) Compile(checkPointer *memory.Memory) Graph {
	// Logic to compile the graph state
	return Graph{
		checkPointer: checkPointer,
		graph:        g.nodes,
		nodeModels:   g.nodeModels,
		currentNode:  string(EdgeStart)}
}

func (g Graph) Invoke(config agents.Config, messages agents.Messages, tools []tools.Tool) string {
	return ""
}
