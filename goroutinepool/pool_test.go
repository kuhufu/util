package goroutinepool

import (
	"log"
	"testing"
	"time"
)

func TestNewEnhancedWorkerPool(t *testing.T) {
	pool := NewWorkerPool(2, 0)

	pool.Run()

	log.Println("Worker pool started")

	for i := 0; i < 2; i++ {
		i := i
		err := pool.AddTask(func() {
			time.Sleep(time.Second * 2)
			log.Println("Task", i, "completed")
		})
		if err != nil {
			return
		}
	}

	pool.Shutdown()
}
