package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel, of type bool, to signal completion
	done := make(chan bool)

	go doWork(done)

	// Wait for the goroutine to signal completion
	<-done
	fmt.Println("Main function continues after goroutine completes.")
}

func doWork(doneCh chan bool) {
	fmt.Println("Doing work...")
	time.Sleep(time.Second * 5)
	fmt.Println("Work done!")
	doneCh <- true
}
