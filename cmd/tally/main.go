// Command tally prints word counts for standard input.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/DanielNetzer/forgecell-trial/tally"
)

func main() {
	in, err := io.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, "tally:", err)
		os.Exit(1)
	}
	counts := tally.Words(string(in))
	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	sort.Strings(words)
	for _, w := range words {
		fmt.Printf("%s %d\n", w, counts[w])
	}
}
