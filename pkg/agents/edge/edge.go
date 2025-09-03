package edge

import "github.com/ochirovch/golanggraph/pkg/agents/state"

type Edge struct {
	From string
	To   string
}

type ConditionalEdgeFunc func(state state.State) (response string)
