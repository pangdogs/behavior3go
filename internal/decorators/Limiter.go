package decorators

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Limiter struct {
	Decorator
	maxLoop int64
}

func (limiter *Limiter) Initialize(setting *BTNodeCfg) {
	limiter.Decorator.Initialize(setting)
	limiter.maxLoop = setting.GetPropertyAsInt64("maxLoop")
	if limiter.maxLoop < 1 {
		panic("maxLoop parameter in Limiter decorator is an obligatory parameter")
	}
}

func (limiter *Limiter) OnTick(tick *Tick) Status {
	if limiter.GetChild() == nil {
		return ERROR
	}

	i := tick.Blackboard.GetInt64("i", tick.GetTree().GetID(), limiter.GetID())

	if i < limiter.maxLoop {
		status := limiter.GetChild().Execute(tick)
		if status == SUCCESS || status == FAILURE {
			tick.Blackboard.Set("i", i+1, tick.GetTree().GetID(), limiter.GetID())
		}
		return status
	}

	return FAILURE
}
