package dal

import (
	"context"
	"log"

	"fmt"
	"github.com/ericfengchao/crawler/Crawler/model"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

var indexName = "crawler"
var typeName = "zhenai"

type saver struct {
	client *elastic.Client
	out    chan interface{}
}

func NewElasticSearchClient() (Saver, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return &saver{
		client: client,
	}, nil
}

func (s *saver) Save(item interface{}) (string, error) {
	if profile, ok := item.(model.Profile); ok {
		resp, err := s.client.Index().Index(indexName).Type(typeName).Id(profile.Name).BodyJson(item).Do(context.Background())
		if err != nil {
			log.Fatalf("Elastic save %+v failed on err %+v\n", item, err)
			return "", err
		}
		//log.Printf("Elastic save resp: %+v\n", resp)
		return resp.Id, nil
	} else {
		return "", errors.New(fmt.Sprintf("Saver: Non Profile object received %+v\n", item))
	}
}

func (s *saver) ItemSaver() chan interface{} {
	if s.out == nil {
		s.out = make(chan interface{})
	}
	go func() {
		itemCount := 0
		for {
			item := <-s.out
			id, err := s.Save(item)
			if err != nil {
				log.Fatalf("Failed saving item %+v on err %+v\n", item, err)
			}
			// successfull save
			itemCount++
			log.Printf("#%d Profile: %s\n", itemCount, id)
		}
	}()
	return s.out
}
