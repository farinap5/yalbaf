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

func TestSelectBasic2(t *testing.T) {
	exp := "SELECT a FROM y where (1=\"a a\" AND 1=2) OR (a=b AND 1=2);"
	l := lexer.Init(exp)
	p := parser.Init(l)
	c := p.Eval()
	
	if c != 29 {
		t.Error("Not the expected token type")
	}
}