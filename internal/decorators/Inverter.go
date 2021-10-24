package decorators

import (
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Inverter struct {
	Decorator
}

func (inverter *Inverter) OnTick(tick *Tick) Status {
	if inverter.GetChild() == nil {
		return ERROR
	}

	status := inverter.GetChild().Execute(tick)

	switch status {
	case SUCCESS:
		status = FAILURE
	case FAILURE:
		status = SUCCESS
	}

	return status
}
