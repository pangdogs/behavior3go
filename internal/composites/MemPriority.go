package composites

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type MemPriority struct {
	Composite
}

func (mp *MemPriority) OnOpen(tick *Tick) {
	tick.GetBlackboard().Set(mp.GetHandle(), "runningChild", int64(0))
}

func (mp *MemPriority) OnTick(tick *Tick) Status {
	child := tick.GetBlackboard().GetInt64(mp.GetHandle(), "runningChild")

	for i := child; i < mp.GetChildCount(); i++ {
		status := mp.GetChild(i).Execute(tick)

		if status != FAILURE {
			if status == RUNNING {
				tick.GetBlackboard().Set(mp.GetHandle(), "runningChild", i)
			}
			return status
		}
	}

	return FAILURE
}
