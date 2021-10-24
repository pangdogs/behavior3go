package behavior3go

import (
	. "github.com/pangdogs/behavior3go/internal/actions"
	. "github.com/pangdogs/behavior3go/internal/composites"
	"github.com/pangdogs/behavior3go/internal/config"
	"github.com/pangdogs/behavior3go/internal/core"
	. "github.com/pangdogs/behavior3go/internal/decorators"
)

type BehaviorTree = core.BehaviorTree

var NewBevTree = core.NewBevTree

type NodeLib = core.NodeLib

func NewNodeLib() *NodeLib {
	nodeLib := core.NewNodeLib()

	// 默认actions
	nodeLib.Register("Error", &Error{})
	nodeLib.Register("Failer", &Failer{})
	nodeLib.Register("Runner", &Runner{})
	nodeLib.Register("Succeeder", &Succeeder{})
	nodeLib.Register("Wait", &Wait{})
	nodeLib.Register("Log", &Log{})

	// 默认composites
	nodeLib.Register("MemPriority", &MemPriority{})
	nodeLib.Register("MemSequence", &MemSequence{})
	nodeLib.Register("Priority", &Priority{})
	nodeLib.Register("Sequence", &Sequence{})

	// 默认decorators
	nodeLib.Register("Inverter", &Inverter{})
	nodeLib.Register("Limiter", &Limiter{})
	nodeLib.Register("MaxTime", &MaxTime{})
	nodeLib.Register("Repeater", &Repeater{})
	nodeLib.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	nodeLib.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})

	return nodeLib
}

type Action = core.Action

type Composite = core.Composite

type Condition = core.Condition

type Decorator = core.Decorator

type SubTree = core.SubTree

var SetSubTreeLoadFunc = core.SetSubTreeLoadFunc

type BTNodeCfg = config.BTNodeCfg

type BTTreeCfg = config.BTTreeCfg

var LoadTreeCfg = config.LoadTreeCfg

type BTProjectCfg = config.BTProjectCfg

var LoadProjectCfg = config.LoadProjectCfg

type RawProjectCfg = config.RawProjectCfg

var LoadRawProjectCfg = config.LoadRawProjectCfg
