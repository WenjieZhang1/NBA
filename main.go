package main

import (
	"crawler/engine"
	"crawler/hupu/parser"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Sch:         &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{URL: "https://nba.hupu.com/stats/players", ParserFunc: parser.ParseTeamlist})
}
