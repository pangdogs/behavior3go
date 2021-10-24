package actions

import (
	b3 "github.com/pangdogs/behavior3go"
)

type Runner struct {
	Action
}

func (r *Runner) OnTick(tick *Tick) b3.Status {
	return b3.RUNNING
}
