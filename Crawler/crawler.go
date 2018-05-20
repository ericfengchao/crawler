package main

import (
	"github.com/ericfengchao/crawler/Crawler/dal"
	"github.com/ericfengchao/crawler/Crawler/engine"
	"github.com/ericfengchao/crawler/Crawler/engine/model"
	"github.com/ericfengchao/crawler/Crawler/fetcher/zhenai/parser"
)

func main() {
	url := "http://city.zhenai.com/"
	e := engine.ConcurrentEngine{
		WorkerCount: 10,
		ItemChan:    dal.ItemSaver(),
	}
	e.Start(model.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
		PageTitle:  "City List Page",
	})
}
