package core

type MemKey struct {
	handle uintptr
	field  string
}

type Blackboard struct {
	tick   Tick
	memory map[MemKey]interface{}
}

func NewBlackboard() *Blackboard {
	b := &Blackboard{}
	b.Initialize()
	return b
}

func (b *Blackboard) Initialize() {
	b.memory = make(map[MemKey]interface{})
}

func (b *Blackboard) GetTick() *Tick {
	return &b.tick
}

func (b *Blackboard) Set(handle uintptr, field string, value interface{}) {
	if b.memory == nil {
		return
	}
	b.memory[MemKey{
		handle: handle,
		field:  field,
	}] = value
}

func (b *Blackboard) Get(handle uintptr, field string) (interface{}, bool) {
	if b.memory == nil {
		return nil, false
	}
	v, ok := b.memory[MemKey{
		handle: handle,
		field:  field,
	}]
	return v, ok
}

func (b *Blackboard) Remove(handle uintptr, field string) {
	if b.memory == nil {
		return
	}
	delete(b.memory, MemKey{
		handle: handle,
		field:  field,
	})
}

func (b *Blackboard) GetFloat64(handle uintptr, field string) float64 {
	v, ok := b.Get(handle, field)
	if !ok {
		return 0
	}
	return v.(float64)
}

func (b *Blackboard) GetInt64(handle uintptr, field string) int64 {
	v, ok := b.Get(handle, field)
	if !ok {
		return 0
	}
	return v.(int64)
}

func (b *Blackboard) GetBool(handle uintptr, field string) bool {
	v, ok := b.Get(handle, field)
	if !ok {
		return false
	}
	return v.(bool)
}

func (b *Blackboard) GetString(handle uintptr, field string) string {
	v, ok := b.Get(handle, field)
	if !ok {
		return ""
	}
	return v.(string)
}
