package scheduler

import "github.com/ericfengchao/crawler/Crawler/engine/model"

type Scheduler interface {
	Run()
	EnqueueWorker(chan model.Request) error
	EnqueueRequest(model.Request) error
}
