package actions

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Succeeder struct {
	Action
}

func (s *Succeeder) OnTick(tick *Tick) Status {
	return SUCCESS
}
