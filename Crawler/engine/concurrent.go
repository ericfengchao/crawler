package engine

import (
	"errors"
	"log"

	"fmt"

	"github.com/ericfengchao/crawler/Crawler/engine/model"
	"github.com/ericfengchao/crawler/Crawler/fetcher"
	"github.com/ericfengchao/crawler/Crawler/scheduler"
)

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   scheduler.Scheduler

	ItemChan chan interface{}
}

func (e *ConcurrentEngine) Start(seeds ...model.Request) {
	e.Scheduler = scheduler.NewConcurrentScheduler()
	e.Scheduler.Run()
	for _, r := range seeds {
		e.Scheduler.EnqueueRequest(r)
	}

	out := make(chan interface{})

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(i, out)
	}

	// unified output/re-iteration
	for o := range out {
		if result, ok := o.(model.ParseResult); ok {
			// if it's user profile, pass it to saving channel
			for _, item := range result.Items {
				go func() {
					e.ItemChan <- item
				}()
			}
			// if it's intermediate parseResult, enqueue it to request channel
			for _, r_ := range result.Requests {
				if isDuplicate(r_.Url) {
					continue
				}
				e.Scheduler.EnqueueRequest(r_)
			}
		} else {
			log.Printf("Got err: %+v\n", o)
		}
	}
}

var uniqueUrlMap = make(map[string]bool)

func isDuplicate(r string) bool {
	if _, ok := uniqueUrlMap[r]; ok {
		// log.Printf("Duplicate url %s\n", r)
		return true
	}
	uniqueUrlMap[r] = true
	return false
}

func (ce *ConcurrentEngine) createWorker(id int, out chan interface{}) {
	log.Printf("Creating worker #%d\n", id)
	in := make(chan model.Request)
	go func() {
		for {
			ce.Scheduler.EnqueueWorker(in)
			r := <-in
			result, err := work(r)
			if err != nil {
				out <- fmt.Errorf("Worker %d encountered error %+v\n", id, err)
			} else {
				out <- result
			}
		}
	}()
}

func work(r model.Request) (model.ParseResult, error) {
	if r.Url == "" {
		log.Println("Invalid Url!", r)
		return model.ParseResult{}, errors.New("Invalid url")
	}
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return model.ParseResult{}, err
	}
	return r.ParserFunc(body, r.PageTitle), nil
}
