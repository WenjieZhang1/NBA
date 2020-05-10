package main

import (
	saver "crawler/distributed/persist"
	"crawler/distributed/persist/rpc_support"
	"github.com/olivere/elastic"
	"log"
)

var host = ":1234"

func main() {
	log.Fatal(serveRpc(host, "player_profile"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpc_support.ServeRpc(host, &saver.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
