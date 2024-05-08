package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var bannerMap map[string]string

func init() {

	bannerMap = make(map[string]string)
	loadBanner("shadow.txt")
	loadBanner("standard.txt")
	loadBanner("thinkertoy.txt")
	// }
}
func loadBanner(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	bannerMap[filename] = strings.Join(lines, "\n")
}

// GetLine retrieves the ASCII art representation for a given character from the specified banner file
func GetLetterArray(char rune, bannerFile string) []string {
	banner, ok := bannerMap[bannerFile]
	if !ok {
		return []string{}
	}
	alphabet := strings.Split(banner, "\n")

	start := (char - 32) * 9
	arr := alphabet[start : start+9]
	return arr
}

func PrintAscii(str string) {
	lines := strings.Split(str, "\n")

	letters := [][]string{}
	for _, line := range lines {

		for _, letter := range line {
			if letter < 32 || letter > 126 {
				fmt.Print("Non ascii")
				return
				{

				}
			}
			arr := GetLetterArray(letter, "standard.txt")
			letters = append(letters, arr)
			if letter == '\n' {
				fmt.Println()
			}
		}
	}

	for i := 1; i < 9; i++ {
		for _, letter := range letters {
			fmt.Printf("%s", letter[i])
		}
		if i < 8 {
			fmt.Println()
		}
	}
}
