package decorators

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
)

type RepeatUntilSuccess struct {
	Decorator
	maxLoop int64
}

func (rus *RepeatUntilSuccess) Initialize(setting *BTNodeCfg) {
	rus.Decorator.Initialize(setting)
	rus.maxLoop = setting.GetPropertyAsInt64("maxLoop")
	if rus.maxLoop < 1 {
		panic("maxLoop parameter in RepeatUntilSuccess decorator is an obligatory parameter")
	}
}

func (rus *RepeatUntilSuccess) OnOpen(tick *Tick) {
	tick.GetBlackboard().Set(rus.GetHandle(), "i", int64(0))
}

func (rus *RepeatUntilSuccess) OnTick(tick *Tick) Status {
	if rus.GetChild() == nil {
		return ERROR
	}

	i := tick.GetBlackboard().GetInt64(rus.GetHandle(), "i")
	status := ERROR

	for rus.maxLoop < 0 || i < rus.maxLoop {
		status = rus.GetChild().Execute(tick)
		if status == FAILURE {
			i++
		} else {
			break
		}
	}

	tick.GetBlackboard().Set(rus.GetHandle(), "i", i)
	return status
}
