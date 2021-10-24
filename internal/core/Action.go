package core

type IAction interface {
	Node
}

type Action struct {
	BaseNode
	BaseWorker
}

func (bn *Action) GetCategory() Category {
	return ACTION
}
