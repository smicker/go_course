package main

import (
	"fmt"
	"runtime/debug"
)

func catchCrash() {
	// If recover() is called, the panic is caught and it is not propagated up to main.
	// This means that only the crashing function ends after defer and the rest of the
	// program continues.
	// If recover() is not called the panic will cause the whole program to exit after
	// the defer function is done. The error would then be reported on stderr.
	if r := recover(); r != nil {
		stack := debug.Stack()
		fmt.Printf("Error: %s\n", r)          // Prints the error
		fmt.Printf("Stacktrace: \n%s", stack) // Prints the stacktrace
	} else {
		fmt.Println("Normal shutdown")
	}
}

func test() {
	defer catchCrash()
	fmt.Println("Doing a panic...")
	panic("Program crashed!!!")
}

func main() {
	test()
	fmt.Println("After crash!!")
}
