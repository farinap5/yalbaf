package main

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
	"github.com/farinap5/yalbaf/internal/parser"
)

func main() {
	exp := "SELECT 1,1,1, FROM y where a=1;"
	l := lexer.Init(exp)
	p := parser.Init(l)
	c := p.Eval()
	fmt.Println(c)
}
