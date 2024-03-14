package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker received stop signal, cleaning up!!")
			return
		default:
			fmt.Println("Worker is doing some work...")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)

	// Simulate some work on the main thread
	time.Sleep(time.Second * 5)
	cancel() // Signal the goroutine to stop

	wg.Wait()
	fmt.Println("All workers have finished")
}
