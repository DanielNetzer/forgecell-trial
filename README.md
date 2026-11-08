# forgecell-trial

A tiny scratch project for trying Forgecell v0.3.0 end to end on a repository outside Forgecell itself. `tally` counts words from standard input.

```sh
echo "Go go lab" | go run ./cmd/tally
```

By default every word is printed with its count, in alphabetical order. Pass `-top N` to print only the `N` most frequent words, highest count first, with ties broken alphabetically:

```sh
echo "pear apple pear banana apple cherry pear" | go run ./cmd/tally -top 3
```

`N` must be a positive integer; any other value prints a usage error to standard error and exits with status 2.
