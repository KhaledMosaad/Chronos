package scheduler

import (
	"context"
	"fmt"

	"github.com/KhaledMosaad/Chronos/internal/task"
)

type Scheduler struct {
	tasks    chan task.Taskable
	workers  int
	stopChan chan struct{}
}

func NewSchedular(workers int, taskBuffer int) *Scheduler {
	s := &Scheduler{
		tasks:    make(chan task.Taskable, taskBuffer),
		workers:  workers,
		stopChan: make(chan struct{}),
	}

	s.Start()
	return s
}

func (s *Scheduler) Start() {
	ctx := context.Background()
	for i := 0; i < s.workers; i++ {
		go func() {
			for { // loop until take a task or stopChan invoked
				select {
				case t := <-s.tasks:
					t.Execute(ctx)
				case <-s.stopChan:
					return
				}
			}
		}()
	}
}

func (s *Scheduler) Submit(task task.Taskable) {
	s.tasks <- task
}

func (s *Scheduler) Stop() {
	close(s.stopChan)
	fmt.Println("Scheduler stopped")
}
