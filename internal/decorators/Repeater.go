package decorators

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Repeater struct {
	Decorator
	maxLoop int64
}

func (repeater *Repeater) Initialize(setting *BTNodeCfg) {
	repeater.Decorator.Initialize(setting)
	repeater.maxLoop = setting.GetPropertyAsInt64("maxLoop")
	if repeater.maxLoop < 1 {
		panic("maxLoop parameter in Repeater decorator is an obligatory parameter")
	}
}

func (repeater *Repeater) OnOpen(tick *Tick) {
	tick.GetBlackboard().Set(tick.GetStack(), "i", int64(0))
}

func (repeater *Repeater) OnTick(tick *Tick) Status {
	if repeater.GetChild() == nil {
		return ERROR
	}

	i := tick.GetBlackboard().GetInt64(tick.GetStack(), "i")
	status := SUCCESS

	for repeater.maxLoop < 0 || i < repeater.maxLoop {
		status = repeater.GetChild().Execute(tick)
		if status == SUCCESS || status == FAILURE {
			i++
		} else {
			break
		}
	}

	tick.GetBlackboard().Set(tick.GetStack(), "i", i)
	return status
}
