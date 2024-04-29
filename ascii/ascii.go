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
func GetLine(char rune, bannerFile string) string {
	banner, ok := bannerMap[bannerFile]
	if !ok {
		return ""
	}

	asciiArt := ""
	for _, line := range strings.Split(banner, "\n") {
		if len(line) == 0 {
			continue
		}
		if line[0] == byte(char) {
			asciiArt = line[1:]
			break
		}
	}
	return asciiArt
}
