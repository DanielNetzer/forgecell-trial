# forgecell-trial

A tiny scratch project for trying Forgecell v0.3.0 end to end on a repository outside Forgecell itself.

## tally

`tally` reads all of standard input, splits the text into words on whitespace, lower-cases each word, counts how often each appears, and prints one line per word as `<word> <count>`, sorted by word. Punctuation stays attached to words because splitting is whitespace-only.

### Install

```sh
go install github.com/DanielNetzer/forgecell-trial/cmd/tally@latest
```

### Usage

```sh
echo "Go go lab" | tally
```

```
go 2
lab 1
```

```sh
tally < notes.txt
```

Redirect a file into `tally` the same way; output is one sorted `<word> <count>` line per distinct lower-cased word.

During development you can also run without installing:

```sh
echo "Go go lab" | go run ./cmd/tally
```

### Tests

```sh
go vet ./...
go test ./...
```
