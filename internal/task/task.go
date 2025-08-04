package task

import (
	"context"
	"fmt"
	"time"
)

type Taskable interface {
	Execute(ctx context.Context) error
}

type CrawlTask struct {
	ID       string
	Priority int
	Timeout  time.Duration
	Params   map[string]any
}

func (t CrawlTask) Execute(ctx context.Context) error {
	fmt.Println("Executing task ", t.ID)
	fmt.Printf("Task: %s Done\n", t.ID)
	return nil
}
