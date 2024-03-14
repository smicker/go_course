package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(name string, timer chan int) {
	// Create a new random source with a seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random number between 1 and 10
	rnd := r.Intn(10) + 1

	time.Sleep(time.Second * time.Duration(rnd))
	timer <- rnd
}

func main() {
	a, b, c := make(chan int), make(chan int), make(chan int)

	go worker("Analytics", a)
	go worker("ACS", b)
	go worker("IT", c)

	// Only wait for the first to finish, or timeout
	select {
	case worktime1 := <-a:
		fmt.Printf("Analytics finished in %ds\n", worktime1)
	case worktime2 := <-b:
		fmt.Printf("ACS finished in %ds\n", worktime2)
	case worktime3 := <-b:
		fmt.Printf("IT finished in %ds\n", worktime3)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout reached, no one finished in time.")
	}
}
