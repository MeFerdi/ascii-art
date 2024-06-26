package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/ascii" // Import the ascii package from the local directory
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go <input_string> [banner_style]")
		return
	}
	str := os.Args[1]
	bannerStyle := "standard.txt" // Default banner style
	if len(os.Args) == 3 {
		if !strings.HasSuffix(os.Args[2], ".txt") {
			bannerStyle = os.Args[2] + ".txt"
		} else {
			bannerStyle = os.Args[2]
		}
	}
	if str == "" {
		return
	}

	modifiedStr, hasError := ascii.SpecialCharacters(str)
	if hasError {
		fmt.Println(modifiedStr)
		return
	}

	str = modifiedStr
	str = strings.ReplaceAll(str, "\n", "\\n")
	lines := strings.Split(str, "\\n")

	newlineCount := 0
	for _, line := range lines {
		if line == "" {
			newlineCount++
			if newlineCount > 0 {
				fmt.Println()
			}
		} else {
			ascii.PrintAscii(line, bannerStyle) // Call the PrintAscii function from the ascii package to print ASCII art
			fmt.Println()
		}
	}
}
