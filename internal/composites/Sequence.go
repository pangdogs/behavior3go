package composites

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Sequence struct {
	Composite
}

func (s *Sequence) OnTick(tick *Tick) Status {
	for i := int64(0); i < s.GetChildCount(); i++ {
		status := s.GetChild(i).Execute(tick)
		if status != SUCCESS {
			return status
		}
	}
	return SUCCESS
}
