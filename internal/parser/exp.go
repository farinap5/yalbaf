package parser

import (
	"github.com/farinap5/yalbaf/internal/lexer"
)

/*
	Parse
	a=1 AND 1=1
	1 AND (1 OR 1)
	"1" AND 1==1
*/
func (p *Parser) parseExp() int {
	p.parseExpGrp()

	if p.Token.Type == lexer.BOOLOP { // AND OR
		//*r = strings.Contains(*data, s)
		auxt := p.Token.Type // aux type since it is replaced by the nxt func
		if auxt == lexer.STRING {
			p.parseStringExpr()
		}
		p.parserGetToken()
		p.parseExpGrp()

		
	}
	return 0
}

func (p *Parser) parseExpGrp() int {
	if p.Token.Type == lexer.OPENGRP {
		p.parserGetToken()
		p.parseExp()
		
		if p.Token.Type != lexer.CLOSEGRP {
			return 0
		}
		p.parserGetToken()
		return 0
	} else {
		return 0
	}
}