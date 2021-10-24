package actions

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Failer struct {
	Action
}

func (f *Failer) OnTick(tick *Tick) Status {
	return FAILURE
}
