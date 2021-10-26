package core

type SubTree struct {
	Action
}

func (st *SubTree) GetCategory() Category {
	return TREE
}

func (st *SubTree) OnTick(tick *Tick) Status {
	// 使用子树，必须先SetSubTreeLoadFunc
	subBevTree := subTreeLoadFunc(st.GetName())
	if subBevTree == nil {
		return ERROR
	}

	tick.pushSubTreeNode(st)
	ret := subBevTree.GetRoot().Execute(tick)
	tick.popSubTreeNode()

	return ret
}

var subTreeLoadFunc func(string) *BehaviorTree

// SetSubTreeLoadFunc 设置子树加载函数
func SetSubTreeLoadFunc(fun func(string) *BehaviorTree) {
	subTreeLoadFunc = fun
}
