package stategraph

import (
	"errors"

	"github.com/ochirovch/golanggraph/pkg/agents/edge"
	"github.com/ochirovch/golanggraph/pkg/agents/invoker"
	"github.com/ochirovch/golanggraph/pkg/agents/node"
	"github.com/ochirovch/golanggraph/pkg/memory"
)

const (
	EdgeStart = "start"
	EdgeEnd   = "end"
)

type conditionalEdge struct {
	fn      edge.ConditionalEdgeFunc
	pathMap map[string]string
}

type StateGraph struct {
	nodes    []node.Node
	edges    map[string][]string
	branches map[string]conditionalEdge
}

func New() *StateGraph {
	return &StateGraph{
		nodes:    make([]node.Node, 0),
		edges:    make(map[string][]string),
		branches: make(map[string]conditionalEdge),
	}
}

func (g *StateGraph) AddNode(name string, values ...any) {
	var llm invoker.Invoker
	var nodeValue interface{}
	for _, v := range values {
		switch n := v.(type) {
		case invoker.Invoker:
			llm = n
		case node.NodeTool, node.NodeFunc:
			nodeValue = n
		}
	}
	switch n := nodeValue.(type) {
	case node.NodeTool:
		g.nodes = append(g.nodes, node.Node{
			Name: name,
			LLM:  nil,
			Tool: n,
		})
	case node.NodeFunc:
		if llm == nil {
			panic("LLM is required for function nodes")
		}
		g.nodes = append(g.nodes, node.Node{
			Name:     name,
			LLM:      llm,
			Function: n,
		})
	}
}

func (g *StateGraph) AddToolNode(name string, toolNode node.NodeTool) {
	g.nodes = append(g.nodes, node.Node{
		Name: name,
		LLM:  nil,
		Tool: toolNode,
	})
}

func (g *StateGraph) AddFuncNode(name string, llm invoker.Invoker, fn node.NodeFunc) {
	g.nodes = append(g.nodes, node.Node{
		Name:     name,
		LLM:      llm,
		Function: fn,
	})
}

func (g *StateGraph) AddEdge(from, to string) {
	g.edges[from] = append(g.edges[from], to)
}

func (g *StateGraph) AddConditionalEdge(
	from string,
	fn edge.ConditionalEdgeFunc,
	pathMap map[string]string) {

	g.branches[from] = conditionalEdge{
		fn:      fn,
		pathMap: pathMap,
	}
}

func (g *StateGraph) checkGraph() error {
	if len(g.nodes) == 0 {
		return errors.New("no nodes in graph")
	}
	if g.nodes[0].Name != EdgeStart {
		return errors.New("graph must start with a start edge")
	}
	if g.nodes[len(g.nodes)-1].Name != EdgeEnd {
		return errors.New("graph must end with an end edge")
	}

	// Check for missing branches
	for _, node := range g.nodes {
		if _, ok := g.branches[node.Name]; ok {
			continue
		}
		if len(g.edges[node.Name]) == 0 {
			return errors.New("node " + node.Name + " has no outgoing edges")
		}
		for _, edge := range g.edges[node.Name] {
			found := false
			for _, n := range g.nodes {
				if n.Name == edge || edge == string(EdgeEnd) {
					found = true
					break
				}
			}
			if !found {
				return errors.New("edge from " + node.Name + " to " + edge + " is broken")
			}
		}
	}

	return nil
}

func (g *StateGraph) Compile(checkPointer *memory.Memory) (Graph, error) {
	if err := g.checkGraph(); err != nil {
		return Graph{}, err
	}
	return Graph{
		checkPointer: checkPointer,
		nodes:        g.nodes,
		edges:        g.edges,
		currentStep:  0,
		branches:     g.branches,
		currentNode:  g.nodes[0],
	}, nil
}

const (
	errEdgeReached        = "cannot invoke graph: reached end"
	errMultipleEdgesFound = "multiple edges found, cannot determine next node"
	errBrokenGraph        = "cannot invoke graph: broken edge"
)
