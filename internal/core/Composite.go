package core

type IComposite interface {
	Node
	GetChildCount() int64
	GetChild(index int64) Node
	AddChild(child Node)
}

type Composite struct {
	BaseNode
	BaseWorker
	children []Node
}

func (c *Composite) GetCategory() Category {
	return COMPOSITE
}

func (c *Composite) GetChildCount() int64 {
	return int64(len(c.children))
}

func (c *Composite) GetChild(index int64) Node {
	return c.children[index]
}

func (c *Composite) AddChild(child Node) {
	c.children = append(c.children, child)
}
