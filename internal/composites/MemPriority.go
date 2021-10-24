package composites

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type MemPriority struct {
	Composite
}

func (mp *MemPriority) OnOpen(tick *Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), mp.GetID())
}

func (mp *MemPriority) OnTick(tick *Tick) Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), mp.GetID())

	for i := child; i < mp.GetChildCount(); i++ {
		status := mp.GetChild(i).Execute(tick)

		if status != FAILURE {
			if status == RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), mp.GetID())
			}

			return status
		}
	}

	return FAILURE
}
