package main

import (
	"github.com/granddaifuku/blahbrah"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(blahbrah.Analyzer) }
