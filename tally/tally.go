// Package tally counts words in text.
package tally

import (
	"strings"
	"unicode"
)

// Words returns how many times each lower-cased word appears in text.
// Leading and trailing Unicode punctuation is trimmed from each token;
// internal punctuation is kept. Punctuation-only tokens are omitted.
func Words(text string) map[string]int {
	counts := map[string]int{}
	for _, w := range strings.Fields(text) {
		w = strings.TrimFunc(w, unicode.IsPunct)
		if w == "" {
			continue
		}
		counts[strings.ToLower(w)]++
	}
	return counts
}
