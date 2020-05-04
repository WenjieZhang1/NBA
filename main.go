package main

import (
	"crawler/engine"
	"crawler/hupu/parser"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Sch:         &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{URL: "https://nba.hupu.com/stats/players", ParserFunc: parser.ParseTeamlist})
}
