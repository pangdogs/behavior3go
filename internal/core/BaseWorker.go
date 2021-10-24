package core

type Worker interface {
	OnEnter(tick *Tick)
	OnOpen(tick *Tick)
	OnTick(tick *Tick) Status
	OnClose(tick *Tick)
	OnExit(tick *Tick)
}

type BaseWorker struct {
}

func (bw *BaseWorker) OnEnter(tick *Tick) {
}

func (bw *BaseWorker) OnOpen(tick *Tick) {
}

func (bw *BaseWorker) OnTick(tick *Tick) Status {
	return ERROR
}

func (bw *BaseWorker) OnClose(tick *Tick) {
}

func (bw *BaseWorker) OnExit(tick *Tick) {
}
