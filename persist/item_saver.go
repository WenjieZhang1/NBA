package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"github.com/olivere/elastic"
	"log"
)

// ItemSaver to save items
func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			err := SaveItem(client, item, "player_profile")
			if err != nil {
				log.Printf("Item saver error saving item: %v: %v", item, err)
			}

		}
	}()
	return out, nil
}

func SaveItem(client *elastic.Client, item engine.Item, index string) error {
	if item.Type == "" {
		return errors.New("must apply type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
