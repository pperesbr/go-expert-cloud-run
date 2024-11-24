package pkg

import (
	"strings"

	"golang.org/x/text/unicode/norm"
	"unicode"
)

func RemoveAccents(s string) string {

	t := norm.NFD.String(s)

	// Remove all non-spacing marks
	var b strings.Builder
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) { // Mn = Nonspacing_Mark
			continue
		}
		b.WriteRune(r)
	}

	return b.String()
}
