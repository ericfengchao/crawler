package scheduler

import (
	"github.com/ericfengchao/crawler/Crawler/engine/model"
)

type SimpleScheduler struct {
	workerChan chan *model.Request
}

func (s *SimpleScheduler) Submit(r *model.Request) {
	s.workerChan <- r
}

func (s *SimpleScheduler) ConfigScheduler(workerChan chan *model.Request) {
	s.workerChan = workerChan
}
