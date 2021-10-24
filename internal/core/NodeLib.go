package core

import (
	"errors"
	"reflect"
)

// NodeLib 节点库
type NodeLib struct {
	nodeMap map[string]reflect.Type
}

// NewNodeLib 创建节点库
func NewNodeLib() *NodeLib {
	nodeLib := &NodeLib{
		nodeMap: make(map[string]reflect.Type),
	}
	return nodeLib
}

// Register 注册节点
func (lib *NodeLib) Register(name string, c interface{}) {
	lib.nodeMap[name] = reflect.TypeOf(c).Elem()
}

// New 创建节点
func (lib *NodeLib) New(name string) (interface{}, error) {
	if v, ok := lib.nodeMap[name]; ok {
		return reflect.New(v).Interface(), nil
	}
	return nil, errors.New("no found node")
}

// Exist 节点是否存在
func (lib *NodeLib) Exist(name string) bool {
	_, ok := lib.nodeMap[name]
	return ok
}
