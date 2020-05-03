package engine

import (
	"crawler/fetcher"
	"log"
)

// ConcurrentEngine which runs worker in parallel
type ConcurrentEngine struct {
	Sch         Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

// Scheduler to schedule requests
type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

// ReadyNotifier notify worker is ready
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run the requests concurrently
func (engine *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	engine.Sch.Run()
	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(engine.Sch.WorkerChan(), out, engine.Sch)
	}

	for _, r := range seeds {
		engine.Sch.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func() { engine.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			engine.Sch.Submit(request)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher, error fetching url: %s, %v\n", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			// Tell scheduler I am ready
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
