// Command tally prints word counts for standard input.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/DanielNetzer/forgecell-trial/tally"
)

func main() {
	fs := flag.NewFlagSet("tally", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	top := fs.Int("top", 0, "print only the `N` most frequent words (positive integer)")
	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(2)
	}
	topSet := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "top" {
			topSet = true
		}
	})
	if topSet && *top <= 0 {
		fmt.Fprintf(os.Stderr, "tally: invalid value %d for flag -top: must be a positive integer\n", *top)
		fs.Usage()
		os.Exit(2)
	}

	in, err := io.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, "tally:", err)
		os.Exit(1)
	}
	counts := tally.Words(string(in))
	var words []string
	if topSet {
		words = tally.Top(counts, *top)
	} else {
		words = make([]string, 0, len(counts))
		for w := range counts {
			words = append(words, w)
		}
		sort.Strings(words)
	}
	for _, w := range words {
		fmt.Printf("%s %d\n", w, counts[w])
	}
}
