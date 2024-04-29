package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if len(os.Args) < 2{
		fmt.Println("Usage: go run main.go <input_string>")
		return
	}

	str := strings.Join(os.Args[1:], " ")
	if str == "" {
		fmt.Println("Error: Empty string")
		return
	}

	// Write text line by line into result
	for i := 0; i < 8; i++ {
		result := ""
		for _, letter := range str {
			result += ascii.GetLine(letter, "standard.txt")
		}
		fmt.Println(result)
		result = ""
	}
}
