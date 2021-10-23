package actions

import (
	b3 "github.com/pangdogs/behavior3go"
	. "github.com/pangdogs/behavior3go/core"
)

type Runner struct {
	Action
}

func (this *Runner) OnTick(tick *Tick) b3.Status {
	return b3.RUNNING
}
