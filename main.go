package main

import (
	"ascii/ascii"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <input_string> [style]")
		return
	}

	input := args[0]
	style := "standard.txt" // Default style
	if len(args) > 1 {
		style = args[1]
	}

	if input == "" {
		fmt.Println("Error: Input string cannot be empty.")
		return
	}

	asciiArt, err := ascii.GenerateASCII(input, style)
	if err != nil {
		fmt.Printf("Error generating ASCII art: %v\n", err)
		return
	}

	fmt.Println(asciiArt)
}
