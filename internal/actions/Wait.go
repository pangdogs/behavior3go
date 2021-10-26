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
	tick.GetBlackboard().Set(w.GetHandle(), "startTime", startTime)
}

func (w *Wait) OnTick(tick *Tick) Status {
	currTime := time.Now().UnixNano() / 1000000
	startTime := tick.GetBlackboard().GetInt64(w.GetHandle(), "startTime")

	if currTime-startTime > w.endTime {
		return SUCCESS
	}

	return RUNNING
}
