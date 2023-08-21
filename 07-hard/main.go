package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Solution 1
	file := os.Args[1]
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(content))

	// Solution 2
	f, err2 := os.Open(os.Args[1])
	if err2 != nil {
		fmt.Println("Error", err2)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
