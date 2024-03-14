package main

import (
	"fmt"
	"time"
)

func cleanup() {
	fmt.Println("Doing cleanup")
}

func test() {
	defer cleanup()
	fmt.Println("Doing a panic...")
	panic("Program crashed!!!")
}

func main() {
	go test()

	time.Sleep(time.Second)
}
