package main

import (
	"github.com/farinap5/yalbaf/internal/lexer"
	"github.com/farinap5/yalbaf/internal/parser"
)

func main() {
	exp := "SELECT xa,y FROM any;"
	l := lexer.Init(exp)
	p := parser.Init(l)
	p.Eval()
}