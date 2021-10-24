package core

type ICondition interface {
	Node
}

type Condition struct {
	BaseNode
	BaseWorker
}

func (c *Condition) GetCategory() Category {
	return CONDITION
}
