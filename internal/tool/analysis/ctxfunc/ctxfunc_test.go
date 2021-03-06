package ctxfunc_test

import (
	"path/filepath"
	"testing"

	"github.com/xqueries/xdb/internal/tool/analysis/ctxfunc"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	dir, err := filepath.Abs("./testdata")
	if err != nil {
		t.Error(err)
	}
	analysistest.Run(t, dir, ctxfunc.Analyzer, "./...")
}
