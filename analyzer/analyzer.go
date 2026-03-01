package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "myloglinter",
	Doc: "Checks logs, logs must: " +
		"start with the small case, " +
		"to be only alphabet symbols and  in English, " +
		"without any sensitive data",
	Run: run,
}
