package blankendline

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const (
	afterLbrace  = "ineffectual blank line after the left brace"
	beforeRbrace = "ineffectual blank line before the right brace"
)

type checker struct {
	pass    *analysis.Pass
	reports map[int]string
}

func newChecker(p *analysis.Pass) *checker {
	return &checker{
		pass:    p,
		reports: make(map[int]string),
	}
}

func (c *checker) report() {
	for offset, r := range c.reports {
		c.pass.Reportf(token.Pos(offset), r)
	}
}

func (c *checker) line(p token.Pos) int {
	return c.pass.Fset.Position(p).Line
}

func (c *checker) col(p token.Pos) int {
	return c.pass.Fset.Position(p).Column
}

func (c *checker) storeReport(offset int, r string) {
	_, ok := c.reports[offset]
	if ok {
		return
	}

	c.reports[offset] = r
}

func (c *checker) blockStmt(
	block *ast.BlockStmt,
) {
	lbraceLine := c.line(block.Lbrace)
	rbraceLine := c.line(block.Rbrace)

	if len(block.List) == 0 {
		if rbraceLine-lbraceLine > 1 {
			c.storeReport(int(block.Rbrace)-1, beforeRbrace)
		}

		return
	}

	firstLine := c.line(block.List[0].Pos())
	firstCol := c.col(block.List[0].Pos())
	if firstLine-lbraceLine > 1 {
		c.storeReport(int(block.Lbrace)+firstCol, afterLbrace)
	}

	endLine := c.line(block.List[len(block.List)-1].End())
	endCol := c.col(block.List[len(block.List)-1].End())
	fmt.Println("endline: ", endLine, "\nRbrace: ", rbraceLine)
	if rbraceLine-endLine > 1 {
		c.storeReport(int(block.Rbrace)-endCol, beforeRbrace)
	}
}

func (c *checker) compositeLit(
	lit *ast.CompositeLit,
) {
	lbraceLine := c.line(lit.Lbrace)
	rbraceLine := c.line(lit.Rbrace)

	firstLine := c.line(lit.Elts[0].Pos())
	firstCol := c.col(lit.Elts[0].Pos())
	if firstLine-lbraceLine > 1 {
		c.storeReport(int(lit.Lbrace)+firstCol, afterLbrace)
	}

	endLine := c.line(lit.Elts[len(lit.Elts)-1].End())
	endCol := c.col(lit.Elts[len(lit.Elts)-1].End())
	if rbraceLine-endLine > 1 {
		fmt.Println("endLine: ", endLine, "\nRBrace: ", rbraceLine)
		c.storeReport(int(lit.Rbrace)-endCol, beforeRbrace)
	}
}
