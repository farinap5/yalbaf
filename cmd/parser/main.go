package main

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
	"github.com/farinap5/yalbaf/internal/parser"
)

func main() {
	exp := "SELECT \"a a\",a,a FROM y where (1=\"a a\" AND 1=2) OR a=b;"
	l := lexer.Init(exp)
	p := parser.Init(l)
	c := p.Eval()
	fmt.Println(c)
}