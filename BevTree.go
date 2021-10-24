package behavior3go

import (
	. "github.com/pangdogs/behavior3go/config"
	"github.com/pangdogs/behavior3go/internal/core"
)

// NewBevTree 创建行为树
func NewBevTree(config *BTTreeCfg, nodeLib *NodeLib) (*BehaviorTree, error) {
	tree := core.NewBevTree()
	if err := tree.Load(config, nodeLib); err != nil {
		return nil, err
	}
	return tree, nil
}
