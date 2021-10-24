package actions

import (
	"fmt"
	. "github.com/pangdogs/behavior3go/internal/config"
	. "github.com/pangdogs/behavior3go/internal/core"
)

type Log struct {
	Action
	info string
}

func (log *Log) Initialize(setting *BTNodeCfg) {
	log.Action.Initialize(setting)
	log.info = setting.GetPropertyAsString("info")
}

func (log *Log) OnTick(tick *Tick) Status {
	fmt.Println("log:", log.info)
	return SUCCESS
}
