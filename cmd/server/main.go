package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KhaledMosaad/Chronos/internal/scheduler"
	"github.com/KhaledMosaad/Chronos/internal/task"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	s := scheduler.NewSchedular(10, 100)
	defer s.Stop()

	for j := 1; j <= 10; j++ {
		var task task.Taskable = task.CrawlTask{
			ID:       fmt.Sprintf("Task: %d", j),
			Priority: 1,
			Timeout:  1 * time.Second,
			Params:   map[string]any{"urls": []string{"https://www.google.com"}},
		}
		s.Submit(task)
	}
	<-sigs
	fmt.Println("Exiting the app")
}
