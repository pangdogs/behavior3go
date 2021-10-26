package core

type Blackboard struct {
	tick   Tick
	memory map[string]interface{}
}

func NewBlackboard() *Blackboard {
	b := &Blackboard{}
	b.Initialize()
	return b
}

func (b *Blackboard) Initialize() {
	b.memory = make(map[string]interface{})
}

func (b *Blackboard) GetTick() *Tick {
	return &b.tick
}

func (b *Blackboard) Set(stack Stack, field string, value interface{}) {
	b.memory[stack.toString()+field] = value
}

func (b *Blackboard) Get(stack Stack, field string) (interface{}, bool) {
	v, ok := b.memory[stack.toString()+field]
	return v, ok
}

func (b *Blackboard) Remove(stack Stack, field string) {
	delete(b.memory, stack.toString()+field)
}

func (b *Blackboard) GetFloat64(stack Stack, field string) float64 {
	v, ok := b.Get(stack, field)
	if !ok {
		return 0
	}
	return v.(float64)
}

func (b *Blackboard) GetInt64(stack Stack, field string) int64 {
	v, ok := b.Get(stack, field)
	if !ok {
		return 0
	}
	return v.(int64)
}

func (b *Blackboard) GetBool(stack Stack, field string) bool {
	v, ok := b.Get(stack, field)
	if !ok {
		return false
	}
	return v.(bool)
}

func (b *Blackboard) GetString(stack Stack, field string) string {
	v, ok := b.Get(stack, field)
	if !ok {
		return ""
	}
	return v.(string)
}
