package parser

import "github.com/farinap5/yalbaf/internal/lexer"

func (p *Parser) SttmSelect() int {
	auxPts := 0
	p.parserGetToken()

	p.parseColumn()

	if p.Token.Type == lexer.FROM {
		p.parserGetToken()
		if p.Token.Type != lexer.IDENTIFIER {
			// TODO: test if subquery gone right
			auxPts += p.parseTableOrSubquery()
		} else {
			auxPts++
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
		//p.parserGetToken()
		if p.Token.Type == lexer.OPENGRP { // Validate subqueries and expressions

		}
		return 0
	}
	p.parserGetToken()

	

	if p.Token.Type == lexer.COMMA {
		p.parserGetToken()
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


func (p *Parser) parseTableOrSubquery() int {
	if p.Token.Type == lexer.IDENTIFIER {
		p.parserGetToken()
		if p.Token.Type == lexer.DOT {
			p.parserGetToken()
			if p.Token.Type != lexer.IDENTIFIER {
				return 0
			}
			p.parserGetToken()
		}

		if p.Token.Type == lexer.AS {
			p.parserGetToken()
			if p.Token.Type != lexer.IDENTIFIER {
				return 0
			}
			p.parserGetToken()
		}
	}
	
	if p.Token.Type == lexer.OPENGRP {
		pts := p.SttmSelect()
		if p.Token.Type != lexer.CLOSEGRP {
			return 0
		}
		return pts + 1
	}
	return 1
}