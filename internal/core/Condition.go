package core

type ICondition interface {
	IBaseNode
}

type Condition struct {
	BaseNode
	BaseWorker
}

func (c *Condition) Ctor() {
	c.category = CONDITION
}
