# takeleton

To make things easier for me, I have automated the review process.

## Installation

### go get

#### `go version` <= go 1.15

```bash
GO111MODULE=off go get github.com/bushiyama/takelton/cmd/takelton
```

#### `go version` >= go 1.16

```bash
go install github.com/bushiyama/takelton/cmd/takelton@latest
```

### Usage

```bash
# $GOPATH/bin/takeleton
$ go vet -vettool=`which takeleton` ./...
```
