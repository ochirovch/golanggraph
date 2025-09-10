package inmemorysaver

import (
	"github.com/google/uuid"
	"github.com/ochirovch/golanggraph/pkg/agents/message"
	"github.com/ochirovch/golanggraph/pkg/agents/state"
	"github.com/ochirovch/golanggraph/pkg/memory"
)

type stateInfo struct {
	uuid        uuid.UUID
	step        int
	currentNode string
}

type InMemorySaver struct {
	stateInfo map[string]stateInfo
	store     map[uuid.UUID]message.Messages
	storeData map[uuid.UUID]map[string]any
}

func New() memory.Memory {
	return &InMemorySaver{
		stateInfo: make(map[string]stateInfo),
		store:     make(map[uuid.UUID]message.Messages),
		storeData: make(map[uuid.UUID]map[string]any),
	}
}

func (s *InMemorySaver) Save(state state.State) {
	s.stateInfo[state.ThreadID] = stateInfo{
		uuid:        state.UUID,
		step:        state.Step,
		currentNode: state.CurrentNode,
	}
	s.store[state.UUID] = state.Messages
	s.storeData[state.UUID] = state.Data
}

func (s *InMemorySaver) Restore(threadID string) ([]state.State, error) {
	if info, exists := s.stateInfo[threadID]; exists {
		return []state.State{{
			ThreadID:    threadID,
			UUID:        info.uuid,
			Step:        info.step,
			CurrentNode: info.currentNode,
			Messages:    s.store[info.uuid],
			Data:        s.storeData[info.uuid],
		}}, nil
	}
	return []state.State{}, nil
}
