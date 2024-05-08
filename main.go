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
	// if conditions for special tabs
	if strings.Contains(str, "\\a") {
		fmt.Println("Error: Bell Character")
		return
	}
	if strings.Contains(str, "\\v") {
		fmt.Println("Error: Vertical tab character")
		return
	}
	if strings.Contains(str, "\\f") {
		fmt.Println("Error:form feed character")
		return
	}
	if strings.Contains(str, "\\r") {
		fmt.Println("Error: carriage ret character")
		return
	}
	if strings.Contains(str, "\\t") {
		str = strings.ReplaceAll(str, "\\t", "    ")

		
	}
	//Handling backspace tabs
	str = strings.ReplaceAll(str, "\\b", "\b")

	for {
		index := strings.Index(str, "\b")

		if index == -1 {
			break
		}
		if index > 0 {
			str = str[:index-1] + str[index+1:]
		} else {
			str = str[index+1:]
		}
	}
	// Split the input string by newline characters
	str = strings.ReplaceAll(str, "\n", "\\n")

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
