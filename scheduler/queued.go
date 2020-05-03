package scheduler

import "crawler/engine"

// QueuedScheduler which create multiple queues for workers
type QueuedScheduler struct {
	requestChan chan engine.Request
	wokerChan   chan chan engine.Request
}

// Submit request to worker
func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

// WorkerReady tells that a worker is ready to receive a request
func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.wokerChan <- w
}

// WorkerChan to ....
func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// Run scheduler
func (q *QueuedScheduler) Run() {
	q.wokerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.wokerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
