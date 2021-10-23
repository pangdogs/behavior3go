package actions

import (
	b3 "github.com/pangdogs/behavior3go"
	. "github.com/pangdogs/behavior3go/core"
)

type Error struct {
	Action
}

func (e *Error) OnTick(tick *Tick) b3.Status {
	return b3.ERROR
}
