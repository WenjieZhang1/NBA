package main

import (
	"crawler/distributed/persist/rpc_support"
	"crawler/engine"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":8000"
	// start ItemSaverServer
	go serveRpc(host, "test1")

	time.Sleep(5 * time.Second)
	// start ItemSaverClient
	client, err := rpc_support.NewClient(host)
	if err != nil {
		panic(err)
	}
	// call Save
	item := engine.Item{}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "" {
		t.Errorf("result: %s; err: %v", result, err)
	}
}
