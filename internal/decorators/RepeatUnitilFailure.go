package decorators

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
)

type RepeatUntilFailure struct {
	Decorator
	maxLoop int64
}

func (ruf *RepeatUntilFailure) Initialize(setting *BTNodeCfg) {
	ruf.Decorator.Initialize(setting)
	ruf.maxLoop = setting.GetPropertyAsInt64("maxLoop")
	if ruf.maxLoop < 1 {
		panic("maxLoop parameter in RepeatUntilFailure decorator is an obligatory parameter")
	}
}

func (ruf *RepeatUntilFailure) OnOpen(tick *Tick) {
	tick.GetBlackboard().Set(ruf.GetHandle(), "i", int64(0))
}

func (ruf *RepeatUntilFailure) OnTick(tick *Tick) Status {
	if ruf.GetChild() == nil {
		return ERROR
	}

	i := tick.GetBlackboard().GetInt64(ruf.GetHandle(), "i")
	status := ERROR

	for ruf.maxLoop < 0 || i < ruf.maxLoop {
		status = ruf.GetChild().Execute(tick)
		if status == SUCCESS {
			i++
		} else {
			break
		}
	}

	tick.GetBlackboard().Set(ruf.GetHandle(), "i", i)
	return status
}
