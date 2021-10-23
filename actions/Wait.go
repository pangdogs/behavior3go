package actions

import (
	b3 "github.com/pangdogs/behavior3go"
	. "github.com/pangdogs/behavior3go/config"
	. "github.com/pangdogs/behavior3go/core"
	"time"
)

/**
 * Wait a few seconds.
 *
 * @module b3
 * @class Wait
 * @extends Action
**/
type Wait struct {
	Action
	endTime int64
}

/**
 * Initialization method.
 *
 * Settings parameters:
 *
 * - **milliseconds** (*Integer*) Maximum time, in milliseconds, a child
 *                                can execute.
 *
 * @method Initialize
 * @param {Object} settings Object with parameters.
 * @construCtor
**/
func (w *Wait) Initialize(setting *BTNodeCfg) {
	w.Action.Initialize(setting)
	w.endTime = setting.GetPropertyAsInt64("milliseconds")
}

/**
 * Open method.
 * @method open
 * @param {Tick} tick A tick instance.
**/
func (w *Wait) OnOpen(tick *Tick) {
	var startTime int64 = time.Now().UnixNano() / 1000000
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), w.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (w *Wait) OnTick(tick *Tick) b3.Status {
	var currTime int64 = time.Now().UnixNano() / 1000000
	var startTime = tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), w.GetID())

	if currTime-startTime > w.endTime {
		return b3.SUCCESS
	}

	return b3.RUNNING
}
