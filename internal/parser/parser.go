package parser

import (
	"github.com/farinap5/yalbaf/internal/lexer"
)

type Parser struct {
	Lexer 		*lexer.Lex
	Token 		lexer.Token
	TrashHold 	int
	Error 		bool
	count		int
}

/*
	Create new parser
*/
func Init(lexer *lexer.Lex) *Parser {
	p := new(Parser)
	p.Lexer = lexer
	p.TrashHold = 3
	p.count = 0
	return p
}

func (p *Parser) parserGetToken() {
	p.count++
	p.Token = p.Lexer.GetToken()
}

func (p *Parser) Eval() int {
	p.parserGetToken()
	if p.Token.Type != lexer.EOF && p.Token.Type != lexer.ERROR && p.Token.Type != lexer.UNKNOWN {
		p.sttmSeq()
		return p.count
	} else {
		return 0
	}
}

func (p *Parser) sttmSeq() int {
	v := p.sttm()
	p.parserGetToken()
	for p.Token.Type == lexer.DOTCOMMA {
		p.parserGetToken()
		v += p.sttm()
	}
	return v
}

func (p *Parser) sttm() int {
	switch p.Token.Type {
	case lexer.SELECT:
		return p.SttmSelect() + 1
	case lexer.UNION:
		p.parserGetToken()
		return p.SttmSelect() + 1
	}
	return 0
}

func (p *Parser) parseStringExpr() int {
	if p.Token.Type != lexer.STRING {
		return 0
	}

	for {
		if p.Token.Type == lexer.STRING {
			break
		}
		p.parserGetToken()
	}

	return 1
}