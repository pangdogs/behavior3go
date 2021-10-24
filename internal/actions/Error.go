package actions

import (
	b3 "github.com/pangdogs/behavior3go"
)

type Error struct {
	Action
}

func (e *Error) OnTick(tick *Tick) b3.Status {
	return b3.ERROR
}
