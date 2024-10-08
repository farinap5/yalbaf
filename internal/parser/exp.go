package parser

import (
	"github.com/farinap5/yalbaf/internal/lexer"
)

func (p *Parser) parseExp() int {
	p.parseExpGrp()

	if p.Token.Type == lexer.BOOLOP {
		//*r = strings.Contains(*data, s)
		//auxt := p.Token.Type // aux type since it is replaced by the nxt func
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
			//return "", errors.New("no statement to close")
		}
		p.parserGetToken()
		return 0
	} else {
		return 0
	}
}