package decorators

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
)

type MaxTime struct {
	Decorator
	maxTime int64
}

func (mt *MaxTime) Initialize(setting *BTNodeCfg) {
	mt.Decorator.Initialize(setting)
	mt.maxTime = setting.GetPropertyAsInt64("maxTime")
	if mt.maxTime < 1 {
		panic("maxTime parameter in MaxTime decorator is an obligatory parameter")
	}
}

func (mt *MaxTime) OnOpen(tick *Tick) {
	startTime := tick.GetNowTime()
	tick.GetBlackboard().Set(mt.GetHandle(), "startTime", startTime)
}

func (mt *MaxTime) OnTick(tick *Tick) Status {
	if mt.GetChild() == nil {
		return ERROR
	}

	currTime := tick.GetNowTime()
	startTime := tick.GetBlackboard().GetInt64(mt.GetHandle(), "startTime")

	status := mt.GetChild().Execute(tick)

	if currTime-startTime > mt.maxTime {
		return FAILURE
	}

	return status
}
