package core

type IComposite interface {
	Node
	GetChildCount() int
	GetChild(index int) Node
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

func (c *Composite) GetChildCount() int {
	return len(c.children)
}

func (c *Composite) GetChild(index int) Node {
	return c.children[index]
}

func (c *Composite) AddChild(child Node) {
	c.children = append(c.children, child)
}
