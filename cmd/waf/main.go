package main

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
	"github.com/farinap5/yalbaf/internal/parser"
)

func main() {
	arg := "admin'/*"
	query := fmt.Sprintf("SELECT ID FROM users WHERE username='%s' AND passwd='aaaa';",arg)
	l := lexer.Init(query)
	p := parser.Init(l)
	c,b := p.Eval()
	fmt.Println(c,b,query)
}
