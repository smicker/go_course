package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	// Create a new random source with a seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random number between 1 and 10
	rnd := r.Intn(10) + 1

	time.Sleep(time.Second * time.Duration(rnd))
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker("Analytics", &wg)
	wg.Add(1)
	go worker("ACS", &wg)
	wg.Add(1)
	go worker("IT", &wg)

	// Wait for every worker to finish
	wg.Wait()
	fmt.Println("All workers have finished")
}
