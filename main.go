package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii" // Import the ascii package from the local directory
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run main.go <input_string>")
		return
	}

	str := strings.Join(os.Args[1:], " ")
	if str == "" {
		fmt.Println("Error: Empty string")
		return
	}
	if strings.Contains(str, "\n") {
		fmt.Println("Error: Input string contains a newline character")
		return
	}
	// Split the input string by newline characters
	str = strings.ReplaceAll(str, "\n", "\\n")

	lines := strings.Split(str, "\\n")

	for _, line := range lines {

		ascii.PrintAscii(line) // Call the PrintAscii function from the ascii package to print ASCII art
		fmt.Println()
		//  str = strings.ReplaceAll(s"tr, "\n", "\n")
		lines := strings.Split(str, "\\n")

		for _, line := range lines {
			if line == "" {
				fmt.Println()
			} else {
				ascii.PrintAscii(line) // Call the PrintAscii function from the ascii package to print ASCII art
				fmt.Println()
			}
		}
	}
}
