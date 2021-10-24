package core

// 常量定义
const (
	VERSION = "0.0.1" // 版本号

	COMPOSITE_NAME = "composite" // 复合节点名称
	DECORATOR_NAME = "decorator" // 装饰器节点名称
	ACTION_NAME    = "action"    // 行为节点名称
	CONDITION_NAME = "condition" // 条件节点名称
)

// Category 节点类型
type Category int8

const (
	COMPOSITE Category = iota // 复合节点
	DECORATOR                 // 装饰器节点
	ACTION                    // 行为节点
	CONDITION                 // 条件节点
)

// Status 返回值定义
type Status int8

const (
	SUCCESS Status = iota // 成功
	FAILURE               // 失败
	RUNNING               // 正在运行
	ERROR                 // 错误
)
