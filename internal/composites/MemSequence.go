package composites

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type MemSequence struct {
	Composite
}

func (ms *MemSequence) OnOpen(tick *Tick) {
	tick.GetBlackboard().Set(tick.GetStack(), "runningChild", int64(0))
}

func (ms *MemSequence) OnTick(tick *Tick) Status {
	child := tick.GetBlackboard().GetInt64(tick.GetStack(), "runningChild")

	for i := child; i < ms.GetChildCount(); i++ {
		status := ms.GetChild(i).Execute(tick)

		if status != SUCCESS {
			if status == RUNNING {
				tick.GetBlackboard().Set(tick.GetStack(), "runningChild", i)
			}
			return status
		}
	}

	return SUCCESS
}
