# takeleton

To make things easier for me, I have automated the review process.

## Installation

### go get

#### `go version` <= go 1.15

```bash
go get github.com/bushiyama/takeleton/cmd/takeleton
```

#### `go version` >= go 1.16

```bash
go install github.com/bushiyama/takeleton/cmd/takeleton@latest
```

### Usage

```bash
# $GOPATH/bin/takeleton
$ go vet -vettool=`which takeleton` ./...
```
