package ascii

import "strings"

// HandleNonASCII handles non-ASCII characters in the input string
func HandleNonASCII(input string) string {
	return strings.Map(func(r rune) rune {
		if r > 127 {
			return '?'
		}
		return r
	}, input)
}
