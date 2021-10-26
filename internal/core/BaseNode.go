package core

import (
	. "github.com/pangdogs/behavior3go/internal/config"
	"unsafe"
)

type Node interface {
	Initialize(setting *BTNodeCfg)
	GetCategory() Category
	GetID() string
	GetName() string
	GetTitle() string
	SetNode(node Node)
	GetNode() Node
	SetWorker(worker Worker)
	GetWorker() Worker
	GetHandle() uintptr
	setSetting(setting *BTNodeCfg)
	GetSetting() *BTNodeCfg
	Execute(tick *Tick) Status
	_execute(tick *Tick) Status
	_enter(tick *Tick)
	_open(tick *Tick)
	_tick(tick *Tick) Status
	_close(tick *Tick)
	_exit(tick *Tick)
}

type BaseNode struct {
	Node
	Worker
	*BTNodeCfg
}

func (bn *BaseNode) Initialize(setting *BTNodeCfg) {
}

func (bn *BaseNode) GetID() string {
	return bn.ID
}

func (bn *BaseNode) GetName() string {
	return bn.Name
}

func (bn *BaseNode) GetTitle() string {
	return bn.Title
}

func (bn *BaseNode) SetNode(node Node) {
	bn.Node = node
}

func (bn *BaseNode) GetNode() Node {
	return bn.Node
}

func (bn *BaseNode) SetWorker(worker Worker) {
	bn.Worker = worker
}

func (bn *BaseNode) GetWorker() Worker {
	return bn.Worker
}

func (bn *BaseNode) GetHandle() uintptr {
	return uintptr(unsafe.Pointer(bn))
}

func (bn *BaseNode) setSetting(setting *BTNodeCfg) {
	bn.BTNodeCfg = setting
}

func (bn *BaseNode) GetSetting() *BTNodeCfg {
	return bn.BTNodeCfg
}

func (bn *BaseNode) Execute(tick *Tick) Status {
	return bn._execute(tick)
}

func (bn *BaseNode) _execute(tick *Tick) Status {
	// ENTER
	bn._enter(tick)

	// OPEN
	if !tick.GetBlackboard().GetBool(bn.GetHandle(), "isOpen") {
		bn._open(tick)
	}

	// TICK
	var status = bn._tick(tick)

	// CLOSE
	if status != RUNNING {
		bn._close(tick)
	}

	// EXIT
	bn._exit(tick)

	return status
}

func (bn *BaseNode) _enter(tick *Tick) {
	tick.enterNode(bn.GetNode())
	bn.OnEnter(tick)
}

func (bn *BaseNode) _open(tick *Tick) {
	tick.openNode(bn.GetNode())
	tick.GetBlackboard().Set(bn.GetHandle(), "isOpen", true)
	bn.OnOpen(tick)
}

func (bn *BaseNode) _tick(tick *Tick) Status {
	tick.tickNode(bn.GetNode())
	return bn.OnTick(tick)
}

func (bn *BaseNode) _close(tick *Tick) {
	tick.GetBlackboard().Set(bn.GetHandle(), "isOpen", false)
	tick.closeNode(bn.GetNode())
	bn.OnClose(tick)
}

func (bn *BaseNode) _exit(tick *Tick) {
	tick.exitNode(bn.GetNode())
	bn.OnExit(tick)
}
