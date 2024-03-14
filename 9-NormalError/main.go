package main

import (
	"fmt"
	"os"
)

// ReadFile reads a file and returns its content and an error if something goes wrong.
func ReadFile(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return data, nil
}

func main() {
	content, err := ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File content:", string(content))
}
