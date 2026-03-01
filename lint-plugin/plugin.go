package main

import (
	"github.com/timeb30/loglint/analyzer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		analyzer.Analyzer,
	)
}
