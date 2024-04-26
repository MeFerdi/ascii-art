package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	if len(os.Args) == 1 || strings.TrimSpace(os.Args[1]) == "" {
		return
	}

	// Writing arguments in a single string
	str := os.Args[1]
	for _, v := range os.Args[2:] {
		str += " " + v
	}

	// 2. Checking weather str contain "\n" or not ---> executing the ascii-art
	prev := 'a'
	severallines := false
	for _, v := range str {
		if v == 'n' && prev == '\\' {
			severallines = true
		}
		prev = v
	}

	// 3. Writing text line by line into result
	result := ""
	if severallines {
		args := strings.Split(str, "\\n")
		for _, word := range args {
			for i := 0; i < 8; i++ {
				for _, letter := range word {
					if ascii.IsASCII(letter) {
						result += ascii.GetLine(1 + int(letter-' ')*9 + i)
					} else {
						fmt.Printf("Non-ASCII character: %c\n", letter)
					}
				}
				fmt.Println(result)
				result = ""
			}
		}
	} else {
		for i := 0; i < 8; i++ {
			for _, letter := range str {
				if ascii.IsASCII(letter) {
					result += ascii.GetLine(1 + int(letter-' ')*9 + i)
				} else {
					fmt.Printf("Error: Non-ASCII character")
				}
			}
			fmt.Println(result)
			result = ""
		}
	}
}
