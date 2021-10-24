package actions

import (
	b3 "github.com/pangdogs/behavior3go"
)

type Succeeder struct {
	Action
}

func (s *Succeeder) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
