package behavior3go

import (
	"github.com/pangdogs/behavior3go/internal/config"
	"github.com/pangdogs/behavior3go/internal/core"
	"github.com/pangdogs/behavior3go/internal/nodelib"
)

type BehaviorTree = core.BehaviorTree

var NewBevTree = core.NewBevTree

type NodeLib = nodelib.NodeLib

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
