# forgecell-trial

A tiny scratch project for trying Forgecell v0.3.0 end to end on a repository outside Forgecell itself.

## tally

`tally` counts words from standard input. It reads all of stdin, splits on whitespace (`strings.Fields`), lower-cases each word, and prints one `word count` line per distinct word, sorted alphabetically. It takes no flags. Punctuation stays attached to words because splitting is on whitespace only.

## Install

```sh
go install github.com/DanielNetzer/forgecell-trial/cmd/tally@latest
```

## Usage

```sh
echo "Go go lab" | tally
```

```
go 2
lab 1
```

```sh
echo "Hello, hello WORLD!" | tally
```

```
hello 1
hello, 1
world! 1
```

## Tests

Run the tests (and optionally vet) the same way CI does:

```sh
go vet ./...
go test ./...
```
