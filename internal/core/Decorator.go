package core

type IDecorator interface {
	IBaseNode
	SetChild(child IBaseNode)
	GetChild() IBaseNode
}

type Decorator struct {
	BaseNode
	BaseWorker
	child IBaseNode
}

func (d *Decorator) Ctor() {
	d.category = DECORATOR
}

func (d *Decorator) SetChild(child IBaseNode) {
	d.child = child
}

func (d *Decorator) GetChild() IBaseNode {
	return d.child
}
