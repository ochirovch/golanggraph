package memory

import (
	"github.com/ochirovch/golanggraph/pkg/agents/state"
)

type Memory interface {
	Save(state.State)
	Restore(threadID string) ([]state.State, error)
}
