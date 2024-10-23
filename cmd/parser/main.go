package main

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
	"github.com/farinap5/yalbaf/internal/parser"
)

func main() {
	exp := "SELECT a FROM b WHERE i=\"aa aa aa1\";"
	l := lexer.Init(exp)
	p := parser.Init(l)
	c,b := p.Eval()
	fmt.Println(c,b)
}
