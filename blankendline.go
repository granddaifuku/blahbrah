package blankendline

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "blankendline finds the blank lines at the end of the code block"

// Analyzer finds ineffectual blank lines after the left brace and before the right brace.
var Analyzer = &analysis.Analyzer{
	Name: "blankendline",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	c := newChecker(pass)

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BlockStmt)(nil),
		(*ast.CompositeLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.BlockStmt:
			c.blockStmt(n)
		case *ast.CompositeLit:
			c.compositeLit(n)
		}
	})

	c.report()

	return nil, nil
}
