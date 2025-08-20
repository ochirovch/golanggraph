package graphstate

import (
	"github.com/dominikbraun/graph"
)

type EdgePoints string

const (
	EdgeStart EdgePoints = "start"
	EdgeEnd   EdgePoints = "end"
)

type GraphState struct {
	graph graph.Graph[string, string]
}

func New() *GraphState {
	return &GraphState{
		graph: graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles()),
	}

}

func (g *GraphState) AddNode(node string) {
	g.graph.AddVertex(node)
}

func (g *GraphState) AddEdge(from, to string) {
	g.graph.AddEdge(from, to)
}

func (g *GraphState) Compile() Graph {
	// Logic to compile the graph state
	return Graph{
		graph: g.graph,
	}
}
