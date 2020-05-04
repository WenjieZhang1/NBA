package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

// ItemSaver to save items
func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCounter := 0
		for {
			item := <-out
			log.Printf("Got item #%d: %v", itemCounter, item)
			itemCounter++

			_, err := saveItem(item)
			if err != nil {
				log.Printf("Item saver error saving item: %v: %v", item, err)
			}

		}
	}()
	return out
}

func saveItem(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().Index("player_profile").Type("NBA").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
