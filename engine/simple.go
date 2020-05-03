package engine

import (
	"crawler/fetcher"
	"log"
)

//SimpleEngine has No current inside
type SimpleEngine struct {
}

func (simpleEngine *SimpleEngine) worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher, error fetching url: %s, %v\n", r.URL, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}

// Run engine
func (simpleEngine *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := simpleEngine.worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("got item: %v\t", item)
		}
	}
}
