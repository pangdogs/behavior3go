package core

import (
	"fmt"
	. "github.com/pangdogs/behavior3go/internal/config"
)

type BehaviorTree struct {
	*BTTreeCfg
	root  Node
	debug interface{}
}

func NewBevTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Initialize()
	return tree
}

func (bt *BehaviorTree) Initialize() {
	bt.root = nil
	bt.debug = nil
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

func (bt *BehaviorTree) SetDebug(debug interface{}) {
	bt.debug = debug
}

func (bt *BehaviorTree) Load(cfg *BTTreeCfg, nodeLib *NodeLib) error {
	nodes := make(map[string]Node)

	// Create the node list (without connection between them)
	for id, nodeCfg := range cfg.Nodes {
		var node Node

		switch nodeCfg.CategoryTag {
		case TREE_TAG:
			node = new(SubTree)
		default:
			if t, err := nodeLib.New(nodeCfg.Name); err == nil {
				node = t.(Node)
			} else {
				return fmt.Errorf("new node %s failed, %v", id, err)
			}
		}

		node.Initialize(nodeCfg)
		node.SetNode(node.(Node))
		node.SetWorker(node.(Worker))
		nodes[id] = node
	}

	// Connect the nodes
	for id, nodeCfg := range cfg.Nodes {
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

	bt.root = nodes[cfg.Root]

	return nil
}

func (bt *BehaviorTree) Tick(target interface{}, blackboard *Blackboard) Status {
	if blackboard == nil {
		return ERROR
	}

	/* CREATE A TICK OBJECT */
	tick := NewTick()
	tick.debug = bt.debug
	tick.target = target
	tick.Blackboard = blackboard
	tick.tree = bt

	/* TICK NODE */
	var state = bt.root._execute(tick)

	/* CLOSE NODES FROM LAST TICK, IF NEEDED */
	var lastOpenNodes = blackboard._getTreeData(bt.GetID()).OpenNodes
	var currOpenNodes []Node
	currOpenNodes = append(currOpenNodes, tick._openNodes...)

	// does not close if it is still open in bt tick
	start := 0
	for i := 0; i < MinInt(len(lastOpenNodes), len(currOpenNodes)); i++ {
		start = i + 1
		if lastOpenNodes[i] != currOpenNodes[i] {
			break
		}
	}

	// close the nodes
	for i := len(lastOpenNodes) - 1; i >= start; i-- {
		lastOpenNodes[i]._close(tick)
	}

	/* POPULATE BLACKBOARD */
	blackboard._getTreeData(bt.GetID()).OpenNodes = currOpenNodes
	blackboard.SetTree("nodeCount", tick._nodeCount, bt.GetID())

	return state
}
