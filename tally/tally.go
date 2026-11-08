// Package tally counts words in text.
package tally

import (
	"sort"
	"strings"
)

// Words returns how many times each lower-cased word appears in text.
func Words(text string) map[string]int {
	counts := map[string]int{}
	for _, w := range strings.Fields(text) {
		counts[strings.ToLower(w)]++
	}
	return counts
}

// Top returns at most n words from counts, highest count first, with ties
// broken alphabetically. It returns nil when n is not positive.
func Top(counts map[string]int, n int) []string {
	if n <= 0 {
		return nil
	}
	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		if counts[words[i]] != counts[words[j]] {
			return counts[words[i]] > counts[words[j]]
		}
		return words[i] < words[j]
	})
	if n < len(words) {
		words = words[:n]
	}
	return words
}
