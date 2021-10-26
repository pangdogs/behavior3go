package core

import (
	"fmt"
	. "github.com/pangdogs/behavior3go/internal/config"
	"time"
	"unsafe"
)

type BehaviorTree struct {
	*BTTreeCfg
	root Node
}

func NewBevTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Initialize()
	return tree
}

func (bt *BehaviorTree) Initialize() {
	bt.root = nil
}

func (bt *BehaviorTree) GetID() string {
	return bt.ID
}

func (bt *BehaviorTree) GetTitle() string {
	return bt.Title
}

func (bt *BehaviorTree) GetRoot() Node {
	return bt.root
}

func (bt *BehaviorTree) GetSetting() *BTTreeCfg {
	return bt.BTTreeCfg
}

func (bt *BehaviorTree) GetHandle() uintptr {
	return uintptr(unsafe.Pointer(bt))
}

func (bt *BehaviorTree) Load(setting *BTTreeCfg, nodeLib *NodeLib) error {
	nodes := make(map[string]Node)

	// Create the node list (without connection between them)
	for id, nodeCfg := range setting.Nodes {
		var node Node

		switch nodeCfg.CategoryTag {
		case TREE_TAG:
			node = &SubTree{}
		default:
			if t, err := nodeLib.New(nodeCfg.Name); err == nil {
				node = t.(Node)
			} else {
				return fmt.Errorf("new node %s failed, %v", id, err)
			}
		}

		category, ok := CategoryTagToEnum[nodeCfg.CategoryTag]
		if ok {
			if node.GetCategory() != category {
				return fmt.Errorf("new node %s failed, category %s invalid", id, nodeCfg.CategoryTag)
			}
		} else {
			return fmt.Errorf("new node %s failed, category %s not found", id, nodeCfg.CategoryTag)
		}

		node.SetNode(node.(Node))
		node.SetWorker(node.(Worker))
		node.setSetting(nodeCfg)
		node.Initialize(nodeCfg)
		nodes[id] = node
	}

	// Connect the nodes
	for id, nodeCfg := range setting.Nodes {
		node := nodes[id]

		switch node.GetCategory() {
		case COMPOSITE:
			for _, cid := range nodeCfg.Children {
				node.(IComposite).AddChild(nodes[cid])
			}

		case DECORATOR:
			if nodeCfg.Child != "" {
				node.(IDecorator).SetChild(nodes[nodeCfg.Child])
			}
		}
	}

	bt.BTTreeCfg = setting
	bt.root = nodes[setting.Root]

	return nil
}

func (bt *BehaviorTree) Tick(target interface{}, blackboard *Blackboard, enableVT bool, virtualTime time.Duration) Status {
	if blackboard == nil {
		return ERROR
	}

	/* CREATE A TICK OBJECT */
	tick := blackboard.GetTick()
	tick.Initialize(blackboard, bt, target, enableVT, virtualTime)

	/* TICK NODE */
	state := bt.GetRoot().Execute(tick)

	/* CLOSE NODES FROM LAST TICK, IF NEEDED */
	var lastOpenNodes []Node
	v, ok := blackboard.Get(bt.GetHandle(), "openNodes")
	if ok {
		lastOpenNodes = v.([]Node)
	}

	var currOpenNodes []Node
	if count := len(tick.openNodes); count > 0 {
		currOpenNodes = make([]Node, count)
		copy(currOpenNodes, tick.openNodes)
	}

	// does not close if it is still open in bt tick
	start := 0
	for i := 0; i < MinInt(len(lastOpenNodes), len(currOpenNodes)); i++ {
		start = i + 1
		if lastOpenNodes[i].GetHandle() != currOpenNodes[i].GetHandle() {
			break
		}
	}

	// close the nodes
	for i := len(lastOpenNodes) - 1; i >= start; i-- {
		lastOpenNodes[i]._close(tick)
	}

	/* POPULATE BLACKBOARD */
	blackboard.Set(bt.GetHandle(), "openNodes", currOpenNodes)

	return state
}
