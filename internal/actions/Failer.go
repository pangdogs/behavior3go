package actions

import (
	b3 "github.com/pangdogs/behavior3go"
)

type Failer struct {
	Action
}

func (f *Failer) OnTick(tick *Tick) b3.Status {
	return b3.FAILURE
}
