package decorators

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
	"time"
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
	startTime := time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), mt.GetID())
}

func (mt *MaxTime) OnTick(tick *Tick) Status {
	if mt.GetChild() == nil {
		return ERROR
	}

	currTime := time.Now().UnixNano() / 1000000
	startTime := tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), mt.GetID())
	status := mt.GetChild().Execute(tick)

	if currTime-startTime > mt.maxTime {
		return FAILURE
	}

	return status
}
