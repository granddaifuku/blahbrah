package main

import (
	"github.com/granddaifuku/blankendline"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(blankendline.Analyzer) }
