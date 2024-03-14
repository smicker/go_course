package main

import (
	"fmt"
	"time"

	"github.com/timandy/routine"
)

func test() {
	fmt.Println(routine.Goid())
}

func main() {
	fmt.Println(routine.Goid())
	test()
	go test()

	time.Sleep(time.Second)
}
