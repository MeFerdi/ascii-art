package ascii

import "unicode"

func IsASCII(r rune) bool {
	return r <= unicode.MaxASCII && unicode.IsGraphic(r)
}
