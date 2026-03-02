package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/timeb30/loglint/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("myplugin", New)
}

type MySettings struct {
	One   string    `json:"one"`
	Two   []Element `json:"two"`
	Three Element   `json:"three"`
}

type Element struct {
	Name string `json:"name"`
}

type PluginExample struct {
	settings MySettings
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[MySettings](settings)
	if err != nil {
		return nil, err
	}

	return &PluginExample{settings: s}, nil
}

func (f *PluginExample) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "myloglint",
			Doc:  "find todos without an author",
			Run:  analyzer.Run,
		},
	}, nil
}

func (f *PluginExample) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
