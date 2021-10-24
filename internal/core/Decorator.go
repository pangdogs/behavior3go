package core

type IDecorator interface {
	Node
	SetChild(child Node)
	GetChild() Node
}

type Decorator struct {
	BaseNode
	BaseWorker
	child Node
}

func (d *Decorator) GetCategory() Category {
	return DECORATOR
}

func (d *Decorator) SetChild(child Node) {
	d.child = child
}

func (d *Decorator) GetChild() Node {
	return d.child
}
