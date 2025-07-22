package scheduler

import (
	"context"
	"fmt"

	"github.com/KhaledMosaad/Chronos/internal/task"
)

type Scheduler struct {
	tasks    chan task.Tasker
	workers  int
	stopChan chan struct{}
}

func NewSchedular(workers int, taskBuffer int) *Scheduler {
	s := &Scheduler{
		tasks:    make(chan task.Tasker, taskBuffer),
		workers:  workers,
		stopChan: make(chan struct{}),
	}

	s.Start()
	return s
}

func (s *Scheduler) Start() {
	for i := 0; i < s.workers; i++ {
		go func() {
			for { // loop until take a task or stopChan invoked
				select {
				case t := <-s.tasks:
					if ct, ok := t.(task.CrawlTask); ok {
						ctx := context.Background()
						ct.Execute(ctx)
						fmt.Printf("Task: %s Done\n", ct.ID)

					}
				case <-s.stopChan:
					return
				}
			}
		}()
	}
}

func (s *Scheduler) Submit(task task.Tasker) {
	s.tasks <- task
}

func (s *Scheduler) Stop() {
	close(s.stopChan)
	fmt.Println("Scheduler stopped")
}
