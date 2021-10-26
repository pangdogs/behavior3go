package core

type Tick struct {
	blackboard       *Blackboard
	bevTree          *BehaviorTree
	target           interface{}
	openNodes        []Node
	openSubtreeNodes []*SubTree
}

func (t *Tick) Initialize(blackboard *Blackboard, bevTree *BehaviorTree, target interface{}) {
	t.blackboard = blackboard
	t.bevTree = bevTree
	t.target = target
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
