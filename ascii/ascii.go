package ascii

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ReadStyleFromFile reads ASCII art style from a file
func ReadStyleFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// GenerateASCII converts a string to ASCII art using the provided style
func GenerateASCII(input string, style string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input string cannot be empty")
	}

	artStyle, err := ReadStyleFromFile(style)
	if err != nil {
		return "", err
	}

	lines := strings.Split(artStyle, "\n")
	asciiArt := ""
	for _, char := range input {
		if char >= 32 && char <= 126 {
			index := int(char) - 32
			if index >= 0 && index < len(lines) {
				asciiArt += lines[index] + "\n"
			} else {
				asciiArt += "?" + "\n" // Placeholder for undefined characters
			}
		} else {
			asciiArt += "?" + "\n" // Placeholder for non-printable characters
		}
	}

	return asciiArt, nil
}
