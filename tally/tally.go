// Package tally counts words in text.
package tally

import "strings"

// Words returns how many times each lower-cased word appears in text.
func Words(text string) map[string]int {
	counts := map[string]int{}
	for _, w := range strings.Fields(text) {
		counts[strings.ToLower(w)]++
	}
	return counts
}
