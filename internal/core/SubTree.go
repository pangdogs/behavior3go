package core

type SubTree struct {
	Action
}

func (st *SubTree) OnTick(tick *Tick) Status {
	// 使用子树，必须先SetSubTreeLoadFunc
	subTree := subTreeLoadFunc(st.GetName())
	if nil == subTree {
		return ERROR
	}

	tick.pushSubtreeNode(st)
	ret := subTree.GetRoot().Execute(tick)
	tick.popSubtreeNode()

	return ret
}

var subTreeLoadFunc func(string) *BehaviorTree

// SetSubTreeLoadFunc 设置子树加载函数
func SetSubTreeLoadFunc(fun func(string) *BehaviorTree) {
	subTreeLoadFunc = fun
}
