package core

import (
	. "github.com/pangdogs/behavior3go/config"
)

type IBaseWrapper interface {
	_execute(tick *Tick) Status
	_enter(tick *Tick)
	_open(tick *Tick)
	_tick(tick *Tick) Status
	_close(tick *Tick)
	_exit(tick *Tick)
}

type IBaseNode interface {
	IBaseWrapper
	Ctor()
	Initialize(params *BTNodeCfg)
	GetCategory() Category
	GetID() string
	GetName() string
	GetTitle() string
	SetBaseWorker(worker IBaseWorker)
	GetBaseWorker() IBaseWorker
	Execute(tick *Tick) Status
}

/**
 * The BaseNode class is used as super class to all nodes in BehaviorJS. It
 * comprises all common variables and methods that a node must have to
 * execute.
 *
 * **IMPORTANT:** Do not inherit from this class, use `b3.Composite`,
 * `b3.Decorator`, `b3.Action` or `b3.Condition`, instead.
 *
 * The attributes are specially designed to serialization of the node in a
 * JSON format. In special, the `parameters` attribute can be set into the
 * visual editor (thus, in the JSON file), and it will be used as parameter
 * on the node initialization at `BehaviorTree.load`.
 *
 * BaseNode also provide 5 callback methods, which the node implementations
 * can override. They are `enter`, `open`, `tick`, `close` and `exit`. See
 * their documentation to know more. These callbacks are called inside the
 * `_execute` method, which is called in the tree traversal.
 *
 * @module b3
 * @class BaseNode
**/
type BaseNode struct {
	IBaseWorker
	*BTNodeCfg
	category Category
}

func (bn *BaseNode) Ctor() {
}

func (bn *BaseNode) Initialize(params *BTNodeCfg) {
	bn.BTNodeCfg = params
}

func (bn *BaseNode) GetCategory() Category {
	return bn.category
}

func (bn *BaseNode) GetID() string {
	return bn.Id
}

func (bn *BaseNode) GetName() string {
	return bn.Name
}

func (bn *BaseNode) GetTitle() string {
	return bn.Title
}

func (bn *BaseNode) SetBaseWorker(worker IBaseWorker) {
	bn.IBaseWorker = worker
}

func (bn *BaseNode) GetBaseWorker() IBaseWorker {
	return bn.IBaseWorker
}

func (bn *BaseNode) Execute(tick *Tick) Status {
	return bn._execute(tick)
}

/**
 * This is the main method to propagate the tick signal to this node. This
 * method calls all callbacks: `enter`, `open`, `tick`, `close`, and
 * `exit`. It only opens a node if it is not already open. In the same
 * way, this method only close a node if the node  returned a status
 * different of `b3.RUNNING`.
 *
 * @method _execute
 * @param {Tick} tick A tick instance.
 * @return {Constant} The tick state.
 * @protected
**/
func (bn *BaseNode) _execute(tick *Tick) Status {
	// ENTER
	bn._enter(tick)

	// OPEN
	if !tick.Blackboard.GetBool("isOpen", tick.tree.id, bn.GetID()) {
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

/**
 * Wrapper for enter method.
 * @method _enter
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (bn *BaseNode) _enter(tick *Tick) {
	tick._enterNode(bn)
	bn.OnEnter(tick)
}

/**
 * Wrapper for open method.
 * @method _open
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (bn *BaseNode) _open(tick *Tick) {
	tick._openNode(bn)
	tick.Blackboard.Set("isOpen", true, tick.tree.id, bn.GetID())
	bn.OnOpen(tick)
}

/**
 * Wrapper for tick method.
 * @method _tick
 * @param {Tick} tick A tick instance.
 * @return {Constant} A state constant.
 * @protected
**/
func (bn *BaseNode) _tick(tick *Tick) Status {
	tick._tickNode(bn)
	return bn.OnTick(tick)
}

/**
 * Wrapper for close method.
 * @method _close
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (bn *BaseNode) _close(tick *Tick) {
	tick._closeNode(bn)
	tick.Blackboard.Set("isOpen", false, tick.tree.id, bn.GetID())
	bn.OnClose(tick)
}

/**
 * Wrapper for exit method.
 * @method _exit
 * @param {Tick} tick A tick instance.
 * @protected
**/
func (bn *BaseNode) _exit(tick *Tick) {
	tick._exitNode(bn)
	bn.OnExit(tick)
}
