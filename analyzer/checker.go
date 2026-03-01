package analyzer

import (
	"github.com/timeb30/loglint/linter"
	"github.com/timeb30/loglint/linter/rules"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func run(pass *analysis.Pass) (any, error) {
	ltr := linter.NewRuleLinker([]linter.Rule{
		rules.LowerCaseRule{},
		rules.EnglishRule{},
		rules.EmojiSpecialRule{},
		rules.SensitiveRule{},
	})
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			obj := pass.TypesInfo.ObjectOf(selExpr.Sel)
			if obj == nil {
				return true
			}
			pkg := obj.Pkg()
			if pkg == nil {
				return true
			}
			pkgPath := pkg.Path()

			if pkgPath != "log/slog" && pkgPath != "go.uber.org/zap" {
				return true
			}
			methodName := selExpr.Sel.Name

			if methodName != "Info" && methodName != "Debug" && methodName != "Error" && methodName != "Warn" {
				return true
			}
			if len(callExpr.Args) == 0 {
				return true
			}
			firstArg := callExpr.Args[0]
			var paramStr string
			var basicLit *ast.BasicLit
			if binaryExpr, ok := firstArg.(*ast.BinaryExpr); ok {
				firstArg = binaryExpr.X
			}
			if basicLit, ok = firstArg.(*ast.BasicLit); ok {
				if ok && basicLit.Kind == token.STRING {
					paramStr = basicLit.Value
				}
			}
			paramStr = strings.Trim(paramStr, "\"`")
			if paramStr == "" {
				return true
			}
			results := ltr.Run(&paramStr)
			for ruleName, passed := range results {
				if !passed {
					pass.Reportf(basicLit.Pos(), ruleName)
				}
			}
			return true
		},
		)
	}
	return nil, nil
}
