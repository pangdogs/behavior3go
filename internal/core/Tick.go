package core

import "time"

type Tick struct {
	bevTree          *BehaviorTree
	target           interface{}
	blackboard       *Blackboard
	enableVT         bool
	virtualTime      time.Duration
	openNodes        []Node
	openSubtreeNodes []*SubTree
}

func (t *Tick) initialize(bevTree *BehaviorTree, target interface{}, blackboard *Blackboard, enableVT bool, virtualTime time.Duration) {
	t.bevTree = bevTree
	t.target = target
	t.blackboard = blackboard
	t.enableVT = enableVT
	t.virtualTime = virtualTime
	t.openNodes = t.openNodes[:0]
	t.openSubtreeNodes = t.openSubtreeNodes[:0]
}

func (t *Tick) GetBevTree() *BehaviorTree {
	return t.bevTree
}

func (t *Tick) GetTarget() interface{} {
	return t.target
}

func (t *Tick) GetBlackboard() *Blackboard {
	return t.blackboard
}

func (t *Tick) GetEnableVT() bool {
	return t.enableVT
}

func (t *Tick) GetVirtualTime() time.Duration {
	return t.virtualTime
}

func (t *Tick) GetNowTime() int64 {
	if t.enableVT {
		return int64(t.virtualTime / time.Millisecond)
	} else {
		return time.Now().UnixNano() / int64(time.Millisecond)
	}
}

func (t *Tick) enterNode(node Node) {
	t.openNodes = append(t.openNodes, node)
}

func (t *Tick) openNode(node Node) {
}

func (t *Tick) tickNode(node Node) {
}

func (t *Tick) closeNode(node Node) {
	if count := len(t.openNodes); count > 0 {
		t.openNodes = t.openNodes[:count-1]
	}
}

func (t *Tick) exitNode(node Node) {
}

func (t *Tick) pushSubTreeNode(node *SubTree) {
	t.openSubtreeNodes = append(t.openSubtreeNodes, node)
}

func (t *Tick) popSubTreeNode() {
	if count := len(t.openSubtreeNodes); count > 0 {
		t.openSubtreeNodes = t.openSubtreeNodes[:count-1]
	}
}

func (t *Tick) getLastSubTreeNode() *SubTree {
	if count := len(t.openSubtreeNodes); count > 0 {
		return t.openSubtreeNodes[count-1]
	}
	return nil
}
