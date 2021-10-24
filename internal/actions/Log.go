package actions

import (
	"fmt"
	b3 "github.com/pangdogs/behavior3go"
	. "github.com/pangdogs/behavior3go/config"
)

type Log struct {
	Action
	info string
}

func (l *Log) Initialize(setting *BTNodeCfg) {
	l.Action.Initialize(setting)
	l.info = setting.GetPropertyAsString("info")
}

func (l *Log) OnTick(tick *Tick) b3.Status {
	fmt.Println("log:", l.info)
	return b3.SUCCESS
}
