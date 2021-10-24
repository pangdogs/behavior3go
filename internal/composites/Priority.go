package composites

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Priority struct {
	Composite
}

func (p *Priority) OnTick(tick *Tick) Status {
	for i := 0; i < p.GetChildCount(); i++ {
		status := p.GetChild(i).Execute(tick)
		if status != FAILURE {
			return status
		}
	}
	return FAILURE
}
