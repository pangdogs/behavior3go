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

func NewBevTreeEx(setting *BTTreeCfg, nodeLib *NodeLib) (*BehaviorTree, error) {
	bevTree := core.NewBevTree()
	if err := bevTree.Load(setting, nodeLib); err != nil {
		return nil, err
	}
	return bevTree, nil
}

type NodeLib = core.NodeLib

var NewNodeLib = core.NewNodeLib

func NewNodeLibEx() *NodeLib {
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

type Tick = core.Tick

type Blackboard = core.Blackboard

const Global = uintptr(0)

var NewBlackboard = core.NewBlackboard

const (
	COMPOSITE_TAG = core.COMPOSITE_TAG
	DECORATOR_TAG = core.DECORATOR_TAG
	ACTION_TAG    = core.ACTION_TAG
	CONDITION_TAG = core.CONDITION_TAG
	TREE_TAG      = core.TREE_TAG
)

type Category = core.Category

const (
	COMPOSITE = core.COMPOSITE
	DECORATOR = core.DECORATOR
	ACTION    = core.ACTION
	CONDITION = core.CONDITION
	TREE      = core.TREE
)

var CategoryTagToEnum = core.CategoryTagToEnum

var CategoryEnumToTag = core.CategoryEnumToTag

type Status = core.Status

const (
	SUCCESS = core.SUCCESS
	FAILURE = core.FAILURE
	RUNNING = core.RUNNING
	ERROR   = core.ERROR
)

type BTNodeCfg = config.BTNodeCfg

type BTTreeCfg = config.BTTreeCfg

var LoadTreeCfg = config.LoadTreeCfg

type BTProjectCfg = config.BTProjectCfg

var LoadProjectCfg = config.LoadProjectCfg

type RawProjectCfg = config.RawProjectCfg

var LoadRawProjectCfg = config.LoadRawProjectCfg
