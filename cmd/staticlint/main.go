// Package staticlint allows one to lint go source code !
//
// # This linter uses
//
// 1. all SA analyzers from staticcheck lib by honnef.co
//
// 2. One analyzer from each of the following packages by honnef.com:
// /simple
// /stylecheck
// /quickfix
// Exact analyzer names can be passed as envs (check env names and defaults in the config file nearby)
//
// 3. Some standard analyzers from x/tools:
// assign, httpresponse, printf, shadow, structtag
//
// 4. Two open-source analyzers:
// inefassign by gordonklaus Detect ineffectual assignments in Go code. An assignment is ineffectual if the variable assigned is not thereafter used.
// maligned by mdempsky which detects structs that would take less memory if their fields were sorted.
//
// 5. A custom analyzer which ensures there is no os.Exit() calls in the main func of the main package of the linted files
//
// USAGE:
// go run ./cmd/staticlint/main.go ./...
// in the project dir
package main

import (
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"

	"github.com/T-V-N/gourlshortener/cmd/staticlint/analyzer"
	"github.com/T-V-N/gourlshortener/cmd/staticlint/config"
	"github.com/gordonklaus/ineffassign/pkg/ineffassign"

	"github.com/mdempsky/maligned/passes/maligned"

	"honnef.co/go/tools/quickfix"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"
)

func main() {
	cfg, err := config.Init()

	if err != nil {
		panic("Unable to parse linter config")
	}

	checks := []*analysis.Analyzer{

		ineffassign.Analyzer,
		maligned.Analyzer,

		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		assign.Analyzer,
		httpresponse.Analyzer,
		analyzer.ExitAnalyzer,
	}

	for _, v := range simple.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, cfg.SimpleAnalyzerName) {
			checks = append(checks, v.Analyzer)
		}
	}

	for _, v := range stylecheck.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, cfg.StylecheckAnalyzerName) {
			checks = append(checks, v.Analyzer)
		}
	}

	for _, v := range quickfix.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, cfg.QuickfixAnalyzerName) {
			checks = append(checks, v.Analyzer)
		}
	}

	for _, v := range staticcheck.Analyzers {
		if !strings.HasPrefix(v.Analyzer.Name, "SA") {
			checks = append(checks, v.Analyzer)
		}
	}

	multichecker.Main(
		checks...,
	)
}
