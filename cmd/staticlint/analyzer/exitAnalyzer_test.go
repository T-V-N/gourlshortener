package analyzer_test

import (
	"testing"

	"github.com/T-V-N/gourlshortener/cmd/staticlint/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test_ExitAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.ExitAnalyzer, "./...")
}
