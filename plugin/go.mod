module github.com/timeb30/logling/plugin

go 1.25.0

require (
	github.com/golangci/plugin-module-register v0.1.2
	github.com/timeb30/loglint v0.0.0-20260301202035-2af83c6eabf1
	golang.org/x/tools v0.42.0
)

replace github.com/timeb30/loglint => ../
