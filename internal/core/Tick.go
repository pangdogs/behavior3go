package core

type Tick struct {
	tree              *BehaviorTree
	debug             interface{}
	target            interface{}
	Blackboard        *Blackboard
	_openNodes        []Node
	_openSubtreeNodes []*SubTree
	_nodeCount        int
}

func NewTick() *Tick {
	tick := &Tick{}
	tick.Initialize()
	return tick
}

func (t *Tick) Initialize() {
	// set by BehaviorTree
	t.tree = nil
	t.debug = nil
	t.target = nil
	t.Blackboard = nil

	// updated during the tick signal
	t._openNodes = t._openNodes[:0]
	t._openSubtreeNodes = t._openSubtreeNodes[:0]
	t._nodeCount = 0
}

func (t *Tick) GetTree() *BehaviorTree {
	return t.tree
}

func (t *Tick) _enterNode(node Node) {
	t._nodeCount++
	t._openNodes = append(t._openNodes, node)

	// TODO: call debug here
}

func (t *Tick) _openNode(node Node) {
	// TODO: call debug here
}

func (t *Tick) _tickNode(node Node) {
	// TODO: call debug here
}

func (t *Tick) _closeNode(node Node) {
	// TODO: call debug here

	count := len(t._openNodes)
	if count > 0 {
		t._openNodes = t._openNodes[:count-1]
	}
}

func (t *Tick) pushSubtreeNode(node *SubTree) {
	t._openSubtreeNodes = append(t._openSubtreeNodes, node)
}

func (t *Tick) popSubtreeNode() {
	count := len(t._openSubtreeNodes)
	if count > 0 {
		t._openSubtreeNodes = t._openSubtreeNodes[:count-1]
	}
}

func (t *Tick) GetLastSubTree() *SubTree {
	count := len(t._openSubtreeNodes)
	if count > 0 {
		return t._openSubtreeNodes[count-1]
	}
	return nil
}

func (t *Tick) _exitNode(node Node) {
	// TODO: call debug here
}

func (t *Tick) GetTarget() interface{} {
	return t.target
}
