package scheduler

import (
	"errors"

	"github.com/ericfengchao/crawler/Crawler/engine/model"
)

type concurrentScheduler struct {
	requestChan chan *model.Request
	workerChan  chan chan *model.Request
}

func NewConcurrentScheduler() Scheduler {
	c := &concurrentScheduler{}
	c.requestChan = make(chan *model.Request)
	c.workerChan = make(chan chan *model.Request)
	return c
}

func (c *concurrentScheduler) EnqueueRequest(r *model.Request) error {
	if r == nil {
		return errors.New("Scheduler: Nil request NOT Allowed!")
	}

	if c.requestChan == nil {
		return errors.New("Scheduler: Request chan not constructed!")
	}

	go func() {
		c.requestChan <- r
	}()
	return nil
}

func (c *concurrentScheduler) EnqueueWorker(worker chan *model.Request) error {
	if worker == nil {
		return errors.New("Scheduler: Nil worker NOT Allowed!")
	}

	if c.workerChan == nil {
		return errors.New("Scheduler: Worker chan not constructed!")
	}

	c.workerChan <- worker
	return nil
}

func (c *concurrentScheduler) Run() {
	go func() {
		c.schedule()
	}()
}

func (c *concurrentScheduler) schedule() {
	var reqQ []*model.Request
	var workerQ []chan *model.Request
	for {
		var activeReq *model.Request
		var activeWorker chan *model.Request
		if len(reqQ) > 0 && len(workerQ) > 0 {
			activeReq = reqQ[0]
			activeWorker = workerQ[0]
		}
		select {
		case req := <-c.requestChan:
			reqQ = append(reqQ, req)
		case worker := <-c.workerChan:
			workerQ = append(workerQ, worker)
		case activeWorker <- activeReq:
			reqQ = reqQ[1:]
			workerQ = workerQ[1:]
			// do stuff here
		}
	}
}
