package composites

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type MemSequence struct {
	Composite
}

func (ms *MemSequence) OnOpen(tick *Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), ms.GetID())
}

func (ms *MemSequence) OnTick(tick *Tick) Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), ms.GetID())

	for i := child; i < ms.GetChildCount(); i++ {
		status := ms.GetChild(i).Execute(tick)

		if status != SUCCESS {
			if status == RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), ms.GetID())
			}

			return status
		}
	}

	return SUCCESS
}
