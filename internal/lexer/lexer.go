package lexer

import "unicode/utf8"

const (
	ERROR = iota
	EOF
	WS // White Space
	STRING
	NUMBER
	COMMENT
	WILDCARD
	OPERATOR
	
	BOOLOP // Boolean Operator

	// ()
	OPENGRP
	CLOSEGRP
)

type Token struct {
	Type uint
	Data string
}

type Lex struct {
	Code 	string
	Cursor 	uint
	//Token	Token
}


/*
	Init will create a new lexer for a different sentence
*/
func Init(Data string) *Lex {
	lex := new(Lex)
	lex.Code = Data
	lex.Cursor = 0
	return lex
}

func (l *Lex) next() rune {
	d, n := utf8.DecodeRuneInString(l.Code[l.Cursor:])
	if n == 0 {
		return 0
	}
	l.Cursor++
	return d
}

func (l *Lex) scanFullSequence(b func(rune) bool) string {
	initPosition := l.Cursor
	nextSymbol := l.next()
	
	for b(nextSymbol) {
		nextSymbol = l.next()
	}

	return l.Code[initPosition:l.Cursor]
}

func (l *Lex) GetToken() Token {
	var token Token

	if int(l.Cursor) == len(l.Code) {
		token = Token{Data: "", Type: EOF}
	}

	data := l.next()

	switch data {
	// First analyze single rune tokens
	case '=' :
		token = Token{Data: "", Type: BOOLOP}
	default:
		if isLetter(data){
			literal := l.scanFullSequence(isLetter)
			tokenType = LookupToken(literal) // token.Type
			tokenLiteral = literal // token.Data
		} else if (isDigit(data)) {
			literal := l.scanFullSequence(isDigit)
			tokenType = NUMBER // token.Type
			tokenLiteral = literal // token.Data
		}
	}

	return token
}
