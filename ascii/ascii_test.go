package ascii

import (
	"testing"
)

func TestPrintAscii(t *testing.T) {
	// Test with a valid string
	validString := "Hello_World!"
	PrintAscii(validString)

	// Test with a non-ASCII character
	nonAsciiString := "Hello, € World!"
	PrintAscii(nonAsciiString)

	// Test with an empty string
	emptyString := ""
	PrintAscii(emptyString)

	// Test with a string containing only newline characters
	newlineString := "\n\n\n"
	PrintAscii(newlineString)
}
