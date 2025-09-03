package state

import (
	"github.com/google/uuid"
	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/agents/node"
)

type State struct {
	ThreadID    string
	UUID        uuid.UUID
	Step        int
	CurrentNode node.Node
	Messages    agents.Messages
	Data        map[string]any
}
