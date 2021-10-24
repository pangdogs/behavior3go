package actions

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
	"time"
)

type Wait struct {
	Action
	endTime int64
}

func (w *Wait) Initialize(setting *BTNodeCfg) {
	w.Action.Initialize(setting)
	w.endTime = setting.GetPropertyAsInt64("milliseconds")
}

func (w *Wait) OnOpen(tick *Tick) {
	startTime := time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), w.GetID())
}

func (w *Wait) OnTick(tick *Tick) Status {
	currTime := time.Now().UnixNano() / 1000000
	startTime := tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), w.GetID())

	if currTime-startTime > w.endTime {
		return SUCCESS
	}

	return RUNNING
}