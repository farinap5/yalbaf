package parser

import "github.com/farinap5/yalbaf/internal/lexer"

func (p *Parser) SttmSelect() int {
	p.parserGetToken()

	p.parseColumn()

	if p.Token.Type == lexer.FROM {
		p.parserGetToken()
		if p.Token.Type != lexer.IDENTIFIER {
			return 0;
		}
	}


	if p.Token.Type == lexer.WHERE {
		if p.parseWhere() < 0 {
			return 0
		}
	}

	if p.Token.Type == lexer.LIMIT {
		p.parserGetToken()
		if p.Token.Type != lexer.NUMBER {
			return 0
		}
	}


	return 0
}

func (p *Parser) parseColumn() int {
	if p.Token.Type != lexer.IDENTIFIER {
		/*if p.Token.Type == lexer.OPENGRP { // Validate subqueries and expressions.

		}*/
		return 0
	}
	p.parserGetToken()

	

	if p.Token.Type == lexer.COMMA {
		return p.parseColumn()
	}

	return 1
}

func (p *Parser) parseWhere() int {
	p.parserGetToken()

	if p.Token.Type != lexer.IDENTIFIER {
		return 0
	}
	p.parserGetToken()

	switch p.Token.Type {
	case lexer.OPERATOR:
		p.parserGetToken()
	default:
		return 0
	}

	
	if p.Token.Type != lexer.IDENTIFIER && p.Token.Type != lexer.NUMBER {
		return 0
	}
	p.parserGetToken()

	return 1
}