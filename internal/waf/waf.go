package waf

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
	"github.com/farinap5/yalbaf/internal/parser"
)

type Waf struct{}

func New() *Waf {
	w := new(Waf)
	return w
}

func (w *Waf)Test(arg string) bool {
	query := fmt.Sprintf("SELECT a FROM b WHERE i='%s';",arg)
	l := lexer.Init(query)
	p := parser.Init(l)
	c,b := p.Eval()

	if !b || c!= 10 {
		return true
	}
	return false
}