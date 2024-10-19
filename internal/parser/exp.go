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
/*func (p *Parser) parseExp() int {
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
}*/


func (p *Parser) parseExp() int {
	if err := p.parseExpGrp(); err != 0 {
		return err
	}

	// Check for boolean operators (AND, OR)
	for p.Token.Type == lexer.BOOLOP {
		p.parserGetToken()

		// Parse the right-hand side of the expression
		if err := p.parseExpGrp(); err != 0 {
			return err
		}

	}

	return 0
}

// parseExpGrp parses an expression group, which could be a parenthesized sub-expression or a basic comparison.
func (p *Parser) parseExpGrp() int {
	// Check if the expression is grouped with parentheses
	if p.Token.Type == lexer.OPENGRP { // '('
		p.parserGetToken()
		
		// Parse the inner expression
		if err := p.parseExp(); err != 0 {
			return err
		}

		if p.Token.Type != lexer.CLOSEGRP { // ')'
			return 0
		}
		p.parserGetToken()
		return 0
	}

	return p.parseSimpleExp()
}

// a = 1 a > 1
func (p *Parser) parseSimpleExp() int {
	if p.Token.Type != lexer.IDENTIFIER && p.Token.Type != lexer.NUMBER && p.parseStringExpr() {
		return 0
	}
	p.parserGetToken()

	// =, >, <, >=, <=, !=
	if p.Token.Type != lexer.EQUAL && p.Token.Type != lexer.GT && p.Token.Type != lexer.LT && 
		p.Token.Type != lexer.GTE && p.Token.Type != lexer.LTE && p.Token.Type != lexer.NE {
		return 0
	}
	p.parserGetToken()

	if p.Token.Type != lexer.NUMBER && p.Token.Type != lexer.IDENTIFIER && p.parseStringExpr() {
		return 0
	}
	p.parserGetToken()

	return 0
}