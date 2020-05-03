package scheduler

import "crawler/engine"

// SimpleScheduler simple version to schedule requests
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Submit request to worker
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

// WorkerReady tell woker is ready
func (s *SimpleScheduler) WorkerReady(r chan engine.Request) {
}

// WorkerChan to ....
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// Run schedule
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
