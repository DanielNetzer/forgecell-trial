# forgecell-trial

A tiny scratch project for trying Forgecell v0.3.0 end to end on a repository outside Forgecell itself.

`tally` reads standard input, splits text on whitespace, and lowercases each word before counting it. It prints each word and its count, sorted by word.

## Install

Requires Go 1.22 or newer.

```sh
go install github.com/DanielNetzer/forgecell-trial/cmd/tally@latest
```

Make sure the directory containing the installed `tally` executable is on your `PATH`.

## Usage

Use the installed executable:

```sh
echo "Go go lab" | tally
```

Output:

```text
go 2
lab 1
```

Or run directly from the repository root:

```sh
echo "one two two" | go run ./cmd/tally
```

Output:

```text
one 1
two 2
```

## Tests

From the repository root, run:

```sh
go test ./...
```
