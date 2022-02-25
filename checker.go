package blankendline

import (
	"go/ast"
	"go/token"
)

const (
	afterLbrace      = "ineffectual blank line after the left brace"
	beforeRbrace     = "ineffectual blank line before the right brace"
	testAfterLbrace  = "want \"" + afterLbrace + "\"\n"
	testBeforeRbrace = "want \"" + beforeRbrace + "\"\n"
)

type report struct {
	pos int
	msg string
}

type checker struct {
	fset     *token.FileSet
	decls    []ast.Decl
	comments map[int]struct{}
}

func newChecker(
	fset *token.FileSet,
	decls []ast.Decl,
	cg []*ast.CommentGroup,
) *checker {
	// create a map whose key is the line number of comments
	comments := make(map[int]struct{})

	for _, c := range cg {
		if c.Text() == testAfterLbrace || c.Text() == testBeforeRbrace {
			continue
		}
		start := fset.Position(c.Pos()).Line
		end := fset.Position(c.End()).Line
		for l := start; l < end+1; l++ {
			comments[l] = struct{}{}
		}
	}

	return &checker{
		fset:     fset,
		decls:    decls,
		comments: comments,
	}
}

func (c *checker) line(p token.Pos) int {
	return c.fset.Position(p).Line
}

func (c *checker) col(p token.Pos) int {
	return c.fset.Position(p).Column
}

func (c *checker) isComment(l int) bool {
	_, ok := c.comments[l]

	return ok
}

func (c *checker) inspect() []report {
	reports := make([]report, 0)

	for _, d := range c.decls {
		switch d := d.(type) {
		case *ast.FuncDecl:
			b := d.Body
			ast.Inspect(b, func(n ast.Node) bool {
				switch n := n.(type) {
				case *ast.BlockStmt:
					r := c.blockStmt(n)
					if r != nil {
						reports = append(reports, r...)
					}
				case *ast.CompositeLit:
					r := c.compositeLit(n)
					if r != nil {
						reports = append(reports, r...)
					}
				}

				return true
			})
		}
	}

	return reports
}

func (c *checker) blockStmt(
	block *ast.BlockStmt,
) []report {
	lbraceLine := c.line(block.Lbrace)
	rbraceLine := c.line(block.Rbrace)

	if len(block.List) == 0 {
		if rbraceLine-lbraceLine > 1 && !c.isComment(lbraceLine+1) {
			return []report{
				{
					pos: int(block.Rbrace) - 1,
					msg: beforeRbrace,
				},
			}
		}

		return nil
	}

	reports := make([]report, 0)

	firstLine := c.line(block.List[0].Pos())
	firstCol := c.col(block.List[0].Pos())
	if firstLine-lbraceLine > 1 && !c.isComment(lbraceLine+1) {
		r := report{
			pos: int(block.Lbrace) + firstCol,
			msg: afterLbrace,
		}
		reports = append(reports, r)
	}

	endLine := c.line(block.List[len(block.List)-1].End())
	endCol := c.col(block.List[len(block.List)-1].End())
	if rbraceLine-endLine > 1 && !c.isComment(rbraceLine-1) {
		r := report{
			pos: int(block.Rbrace) - endCol,
			msg: beforeRbrace,
		}
		reports = append(reports, r)
	}

	return reports
}

func (c *checker) compositeLit(
	lit *ast.CompositeLit,
) []report {
	lbraceLine := c.line(lit.Lbrace)
	rbraceLine := c.line(lit.Rbrace)

	if len(lit.Elts) == 0 {
		if rbraceLine-lbraceLine > 1 && !c.isComment(lbraceLine+1) {
			return []report{
				{
					pos: int(lit.Rbrace) - 1,
					msg: beforeRbrace,
				},
			}
		}

		return nil
	}

	reports := make([]report, 0)

	firstLine := c.line(lit.Elts[0].Pos())
	firstCol := c.col(lit.Elts[0].Pos())
	if firstLine-lbraceLine > 1 && !c.isComment(lbraceLine+1) {
		r := report{
			pos: int(lit.Lbrace) + firstCol,
			msg: afterLbrace,
		}
		reports = append(reports, r)
	}

	endLine := c.line(lit.Elts[len(lit.Elts)-1].End())
	endCol := c.col(lit.Elts[len(lit.Elts)-1].End())
	if rbraceLine-endLine > 1 && !c.isComment(rbraceLine-1) {
		r := report{
			pos: int(lit.Rbrace) - endCol,
			msg: beforeRbrace,
		}
		reports = append(reports, r)

	}

	return reports
}
