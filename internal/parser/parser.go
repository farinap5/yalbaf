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

func (p *Parser) Eval() (int, bool) {
	p.parserGetToken()
	if p.Token.Type != lexer.EOF && p.Token.Type != lexer.ERROR && p.Token.Type != lexer.UNKNOWN {
		return p.count, p.sttmSeq()
	} else {
		return 0, false
	}
}

func (p *Parser) sttmSeq() bool {
	if !p.sttm() {
		return false
	}
	//p.parserGetToken()
	for p.Token.Type == lexer.DOTCOMMA {
		p.parserGetToken()
		if p.Token.Type == lexer.EOF {
			return true
		}

		if !p.sttm() {
			return false
		}
	}
	return true
}

func (p *Parser) sttm() bool {
	switch p.Token.Type {
	case lexer.SELECT:
		return p.SttmSelect()
	case lexer.UNION:
		p.parserGetToken()
		return p.SttmSelect()
	}
	return false
}

func (p *Parser) parseStringExpr() bool {
	if p.Token.Type != lexer.STRING {
		return false
	}
	aux := p.Token.Type // store the token to use the same for breaking the loop
	p.parserGetToken()

	for {
		if p.Token.Type == lexer.STRING && p.Token.Type == aux {
			break
		}
		if p.Token.Type == lexer.EOF {
			return false
		}
		p.parserGetToken()
	}

	return true
}