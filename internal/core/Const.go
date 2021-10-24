package core

// 节点类型名称
const (
	COMPOSITE_TAG = "composite"
	DECORATOR_TAG = "decorator"
	ACTION_TAG    = "action"
	CONDITION_TAG = "condition"
	TREE_TAG      = "tree"
)

// Category 节点类型
type Category int8

const (
	COMPOSITE Category = iota // 复合节点
	DECORATOR                 // 装饰器节点
	ACTION                    // 行为节点
	CONDITION                 // 条件节点
	TREE                      // 子树节点
)

// Status 返回值定义
type Status int8

const (
	SUCCESS Status = iota // 成功
	FAILURE               // 失败
	RUNNING               // 正在运行
	ERROR                 // 错误
)