package parser

import (
	"github.com/farinap5/yalbaf/internal/lexer"
)

func (p *Parser) SttmSelect() bool {
	p.parserGetToken()

	if p.Token.Type != lexer.EOF {
		if !p.parseColumn() {
			return false
		}
	}

	if p.Token.Type == lexer.FROM {
		p.parserGetToken()
		if p.Token.Type != lexer.IDENTIFIER {
			// TODO: test if subquery gone right
			if !p.parseTableOrSubquery() {
				return false
			}
		}
		p.parserGetToken()
	}


	if p.Token.Type == lexer.WHERE {
		if !p.parseWhere() {
			return false
		}
	}

	if p.Token.Type == lexer.LIMIT {
		p.parserGetToken()
		if p.Token.Type != lexer.NUMBER {
			return false
		}
	}


	if p.Token.Type == lexer.UNION {
		p.parserGetToken()
		if p.Token.Type != lexer.NUMBER {
			return false
		}
	}


	return true
}

func (p *Parser) parseColumn() bool {
	if p.Token.Type != lexer.IDENTIFIER && p.Token.Type != lexer.NUMBER {
		//p.parserGetToken()
		if p.Token.Type == lexer.OPENGRP { // Validate sub-queries and expressions
			if !p.sttm() {
				return false
			}
		}
		if !p.parseStringExpr() {
			return false
		}
	}
	p.parserGetToken()

	

	if p.Token.Type == lexer.COMMA {
		p.parserGetToken()
		return p.parseColumn()
	}

	return true
}

func (p *Parser) parseWhere() bool {
	p.parserGetToken()
	return p.parseExp()

	/*if p.Token.Type != lexer.IDENTIFIER {
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

	return 1*/
}


func (p *Parser) parseTableOrSubquery() bool {
	if p.Token.Type == lexer.IDENTIFIER {
		p.parserGetToken()
		if p.Token.Type == lexer.DOT {
			p.parserGetToken()
			if p.Token.Type != lexer.IDENTIFIER {
				return false
			}
			p.parserGetToken()
		}

		if p.Token.Type == lexer.AS {
			p.parserGetToken()
			if p.Token.Type != lexer.IDENTIFIER {
				return false
			}
			p.parserGetToken()
		}
	}
	

	if p.Token.Type == lexer.OPENGRP {
		if !p.SttmSelect() {
			return false
		}
		if p.Token.Type != lexer.CLOSEGRP {
			return false
		}
	}
	return true
}