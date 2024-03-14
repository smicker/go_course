package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("Doing a panic...")
	panic("Program crashed!!!")
}

func main() {
	go test()

	time.Sleep(time.Second)
}
