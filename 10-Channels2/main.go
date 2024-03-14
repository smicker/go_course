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

	// Wait for every worker to finish
	worktime1, worktime2, worktime3 := <-a, <-b, <-c
	fmt.Printf("Analytics finished in %ds\n", worktime1)
	fmt.Printf("ACS finished in %ds\n", worktime2)
	fmt.Printf("IT finished in %ds\n", worktime3)
}
