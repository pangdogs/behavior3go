package core

import (
	"unsafe"
)

type Stack []byte

func (stack *Stack) ToString() string {
	if len(*stack) <= 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(stack))
}

type Tick struct {
	blackboard       *Blackboard
	bevTree          *BehaviorTree
	target           interface{}
	stack            Stack
	openNodes        []Node
	openSubtreeNodes []*SubTree
}

func (t *Tick) Initialize(blackboard *Blackboard, bevTree *BehaviorTree, target interface{}) {
	t.blackboard = blackboard
	t.bevTree = bevTree
	t.target = target
	t.stack = t.stack[:0]
	t.openNodes = t.openNodes[:0]
	t.openSubtreeNodes = t.openSubtreeNodes[:0]
}

func (t *Tick) GetBlackboard() *Blackboard {
	return t.blackboard
}

func (t *Tick) GetBevTree() *BehaviorTree {
	return t.bevTree
}

func (t *Tick) GetTarget() interface{} {
	return t.target
}

func (t *Tick) GetStack() Stack {
	return t.stack
}

func (t *Tick) enterNode(node Node) {
	handle := node.getHandle()
	t.stack = append(t.stack,
		byte((handle>>56)&uintptr(0xff)),
		byte((handle>>48)&uintptr(0xff)),
		byte((handle>>40)&uintptr(0xff)),
		byte((handle>>32)&uintptr(0xff)),
		byte((handle>>24)&uintptr(0xff)),
		byte((handle>>16)&uintptr(0xff)),
		byte((handle>>8)&uintptr(0xff)),
		byte(handle&uintptr(0xff)),
	)
	t.openNodes = append(t.openNodes, node)
}

func (t *Tick) openNode(node Node) {
}

func (t *Tick) tickNode(node Node) {
}

func (t *Tick) closeNode(node Node) {
	if count := len(t.stack); count >= 8 {
		t.stack = t.stack[:count-8]
	}
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
