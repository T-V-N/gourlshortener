// Package analyzer is an analyzer which implements analysis.Analyzer interface
// it ensures there are no os.Exit calls in the main func of the main package
package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// ExitAnalyzer implements Analyzer interface so as to make it possible to use it in multichecker
var ExitAnalyzer = &analysis.Analyzer{
	Name: "ExitAnalyzer",
	Doc:  "Usage of os.Exit in the main package is banned",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.Package:
				if x.Name != "main" {
					return false
				}
			case *ast.FuncDecl:
				if x.Name.Name != "main" {
					return false
				}

			case *ast.SelectorExpr:
				if x.Sel.Name == "Exit" {
					pass.Reportf(x.Pos(), "Usage of os.Exit in the main package is banned")
				}
			}
			return true
		})
	}

	return nil, nil
}
