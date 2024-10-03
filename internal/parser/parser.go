package parser

import (
	"github.com/farinap5/yalbaf/internal/lexer"
)

type Parser struct {
	Lexer lexer.Lex
	Token lexer.Token
	TrashHold int
}

/*
	Create new parser
*/
func Init(lexer lexer.Lex) *Parser {
	p := new(Parser)
	p.Lexer = lexer
	p.TrashHold = 3
	return p
}

func (p *Parser) parserGetToken() {
	p.Token = p.Lexer.GetToken()
}

func (p *Parser) Eval() int {
	p.parserGetToken()
	if p.Token.Type != lexer.EOF || p.Token.Type != lexer.ERROR || p.Token.Type != lexer.UNKNOWN {
		return p.sttmSeq()
	} else {
		return 0
	}
}

func (p *Parser) sttmSeq() int {
	v := p.sttm()
	for p.Token.Type == lexer.DOTCOMMA {
		v += p.sttm()
	}
	return v
}

func (p *Parser) sttm() int {
	switch p.Token.Type {
	case lexer.SELECT:
		return p.SttmSelect()
	}
	return 0
}

