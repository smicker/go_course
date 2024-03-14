package main

import (
	"fmt"
	"time"
)

func cleanup() {
	fmt.Println("Doing cleanup")
}

func test() {
	fmt.Println("Doing a panic...")
	panic("Program crashed!!!")
}

func main() {
	defer cleanup()
	go test()

	time.Sleep(time.Second)
}
