package dal

import (
	"context"
	"log"

	"github.com/ericfengchao/crawler/Crawler/model"
	"github.com/olivere/elastic"
)

var indexName = "crawler"
var typeName = "zhenai"

func save(item interface{}) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Elastic init failed on %v\n", err)
	}

	if profile, ok := item.(model.Profile); ok {
		resp, err := client.Index().Index(indexName).Type(typeName).Id(profile.Name).BodyJson(item).Do(context.Background())
		if err != nil {
			log.Fatalf("Elastic save %+v failed on err %+v\n", item, err)
			return
		}
		log.Printf("Elastic save resp: %+v\n", resp)
	}
}

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			// save(item)
			log.Printf("#%d Profile: %+v", itemCount, item)
		}
	}()
	return out
}
