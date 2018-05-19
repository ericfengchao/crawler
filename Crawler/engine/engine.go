package engine

import (
	"errors"
	"log"

	"github.com/ericfengchao/crawler/Crawler/engine/model"
	"github.com/ericfengchao/crawler/Crawler/fetcher"
)

func Run(seeds ...model.Request) {
	requestChan := make(chan model.Request, 10)
	for _, r := range seeds {
		requestChan <- r
	}
	parsingResult := make(chan model.ParseResult, 10)
	go func() {
		consume(parsingResult, requestChan)
	}()
	for r := range requestChan {
		go func(r model.Request) {
			producer(r, parsingResult)
		}(r)
	}
}

func producer(r model.Request, parseResultChan chan model.ParseResult) error {
	if r.Url == "" {
		log.Println("Invalid Url!", r)
		return errors.New("Invalid url")
	}
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url: %s, err: %v\n", r.Url, err)
		return err
	}
	if newR := r.ParserFunc(body); newR != nil {
		parseResultChan <- *newR
	}
	return nil
}

func consume(parseResultChan chan model.ParseResult, requestChan chan model.Request) error {
	for pr := range parseResultChan {
		for _, r := range pr.Requests {
			requestChan <- r
		}
		//deal with the content
		for _, v := range pr.Items {
			log.Printf("Got item %v\n", v)
		}
	}
	return nil
}
