package state

import (
	"github.com/google/uuid"
	"github.com/ochirovch/golanggraph/pkg/agents/message"
)

type State struct {
	ThreadID    string
	UUID        uuid.UUID
	Step        int
	CurrentNode string
	Messages    message.Messages
	Data        map[string]any
}
