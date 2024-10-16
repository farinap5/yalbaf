package parser_test

import (
	"testing"

	"github.com/farinap5/yalbaf/internal/lexer"
	"github.com/farinap5/yalbaf/internal/parser"
)

func TestSelectBasic(t *testing.T) {
	exp := "SELECT 6,a,b FROM any;"
	l := lexer.Init(exp)
	p := parser.Init(l)
	c := p.Eval()
	if c != 10 {
		t.Error("Not the expected token type")
	}
}