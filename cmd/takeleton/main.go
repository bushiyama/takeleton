package main

import (
	"github.com/bushiyama/takeleton"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(takeleton.Analyzer) }
