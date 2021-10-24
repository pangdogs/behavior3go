package core

type IComposite interface {
	IBaseNode
	GetChildCount() int
	GetChild(index int) IBaseNode
	AddChild(child IBaseNode)
}

type Composite struct {
	BaseNode
	BaseWorker
	children []IBaseNode
}

func (c *Composite) Ctor() {
	c.category = COMPOSITE
}

func (c *Composite) GetChildCount() int {
	return len(c.children)
}

func (c *Composite) GetChild(index int) IBaseNode {
	return c.children[index]
}

func (c *Composite) AddChild(child IBaseNode) {
	c.children = append(c.children, child)
}
