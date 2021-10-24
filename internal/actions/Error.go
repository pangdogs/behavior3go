package actions

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Error struct {
	Action
}

func (e *Error) OnTick(tick *Tick) Status {
	return ERROR
}
