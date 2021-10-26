package core

import (
	"strings"
)

type Stack []uintptr

func (stack *Stack) pushHandle(handle uintptr) {
	*stack = append(*stack, handle)
}

func (stack *Stack) popHandle() {
	if count := len(*stack); count > 0 {
		*stack = (*stack)[:count-1]
	}
}

func (stack *Stack) toString() string {
	b := strings.Builder{}

	for _, handle := range *stack {
		b.WriteByte(byte((handle >> 56) & uintptr(0xff)))
		b.WriteByte(byte((handle >> 48) & uintptr(0xff)))
		b.WriteByte(byte((handle >> 40) & uintptr(0xff)))
		b.WriteByte(byte((handle >> 32) & uintptr(0xff)))
		b.WriteByte(byte((handle >> 24) & uintptr(0xff)))
		b.WriteByte(byte((handle >> 16) & uintptr(0xff)))
		b.WriteByte(byte((handle >> 8) & uintptr(0xff)))
		b.WriteByte(byte(handle & uintptr(0xff)))
	}

	return b.String()
}

func (stack *Stack) Copy() Stack {
	t := make(Stack, len(*stack))
	copy(t, *stack)
	return t
}

func (stack *Stack) Equal(other Stack) bool {
	if len(*stack) != len(other) {
		return false
	}

	for i, b := range *stack {
		if b != other[i] {
			return false
		}
	}

	return true
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

func (t *Tick) GetLastSubTreeStack() Stack {
	subTreeNode := t.getLastSubTreeNode()
	if subTreeNode == nil {
		return Stack{}
	}

	var stack Stack
	stack.pushHandle(subTreeNode.getHandle())

	return stack
}

func (t *Tick) enterNode(node Node) {
	t.stack.pushHandle(node.getHandle())
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
	t.stack.popHandle()
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
