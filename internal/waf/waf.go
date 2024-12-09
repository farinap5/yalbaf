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

func (w *Waf)TestStr(arg string) (int, bool) {
	query := fmt.Sprintf("SELECT a FROM b WHERE i='%s';",arg)
	l := lexer.Init(query)
	p := parser.Init(l)
	c,b := p.Eval()

	if !b || c!= 10 {
		return c-9, true
	}
	return 0, false
}

func (w *Waf)TestInt(arg string) (int, bool) {
	query := fmt.Sprintf("SELECT a FROM b WHERE i=%s;",arg)
	l := lexer.Init(query)
	p := parser.Init(l)
	c,b := p.Eval()

	if !b || c!= 10 {
		return c-9, true
	}
	return 0, false
}