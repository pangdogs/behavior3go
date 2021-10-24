package actions

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Runner struct {
	Action
}

func (r *Runner) OnTick(tick *Tick) Status {
	return RUNNING
}
