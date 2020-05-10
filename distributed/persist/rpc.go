package saver

import (
	"crawler/engine"
	"crawler/persist"
	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (saver *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.SaveItem(saver.Client, item, saver.Index)
	if err == nil {
		*result = "ok"
		return nil
	}
	return err
}
