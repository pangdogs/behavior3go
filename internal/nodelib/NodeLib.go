package nodelib

import (
	"errors"
	. "github.com/pangdogs/behavior3go/internal/actions"
	. "github.com/pangdogs/behavior3go/internal/composites"
	. "github.com/pangdogs/behavior3go/internal/decorators"
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

	// 默认actions
	nodeLib.Register("Error", &Error{})
	nodeLib.Register("Failer", &Failer{})
	nodeLib.Register("Runner", &Runner{})
	nodeLib.Register("Succeeder", &Succeeder{})
	nodeLib.Register("Wait", &Wait{})
	nodeLib.Register("Log", &Log{})

	// 默认composites
	nodeLib.Register("MemPriority", &MemPriority{})
	nodeLib.Register("MemSequence", &MemSequence{})
	nodeLib.Register("Priority", &Priority{})
	nodeLib.Register("Sequence", &Sequence{})

	// 默认decorators
	nodeLib.Register("Inverter", &Inverter{})
	nodeLib.Register("Limiter", &Limiter{})
	nodeLib.Register("MaxTime", &MaxTime{})
	nodeLib.Register("Repeater", &Repeater{})
	nodeLib.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	nodeLib.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})

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
