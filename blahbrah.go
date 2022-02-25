package blahbrah

import (
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const doc = "blahbrah finds ineffectual blank lines after the left brace and before the right brace"

// Analyzer finds ineffectual blank lines after the left brace and before the right brace.
var Analyzer = &analysis.Analyzer{
	Name: "blahbrah",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		c := newChecker(pass.Fset, f.Decls, f.Comments)
		reports := c.inspect()
		for _, r := range reports {
			pass.Reportf(token.Pos(r.pos), r.msg)
		}
	}

	return nil, nil
}
