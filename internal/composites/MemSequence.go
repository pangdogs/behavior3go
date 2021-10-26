package composites

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type MemSequence struct {
	Composite
}

func (ms *MemSequence) OnOpen(tick *Tick) {
	tick.GetBlackboard().Set(ms.GetHandle(), "runningChild", int64(0))
}

func (ms *MemSequence) OnTick(tick *Tick) Status {
	child := tick.GetBlackboard().GetInt64(ms.GetHandle(), "runningChild")

	for i := child; i < ms.GetChildCount(); i++ {
		status := ms.GetChild(i).Execute(tick)

		if status != SUCCESS {
			if status == RUNNING {
				tick.GetBlackboard().Set(ms.GetHandle(), "runningChild", i)
			}
			return status
		}
	}

	return SUCCESS
}
