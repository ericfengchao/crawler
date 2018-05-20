package main

import (
	"github.com/ericfengchao/crawler/Crawler/dal"
	"github.com/ericfengchao/crawler/Crawler/engine"
	"github.com/ericfengchao/crawler/Crawler/engine/model"
	"github.com/ericfengchao/crawler/Crawler/fetcher/zhenai/parser"
)

func main() {
	url := "http://city.zhenai.com/"
	saver, err := dal.NewElasticSearchClient()
	if err != nil {
		// elastic search client failed
		panic(err)
	}
	e := engine.ConcurrentEngine{
		WorkerCount: 10,
		ItemChan:    saver.ItemSaver(),
	}
	e.Start(model.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
		PageTitle:  "City List Page",
	})
}
