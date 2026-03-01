package main

import (
	"github.com/timeb30/loglint/analyzer"
	"golang.org/x/tools/go/analysis"
)

func New() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}
