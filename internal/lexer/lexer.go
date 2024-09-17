package lexer

import (
	"unicode"
	"unicode/utf8"
)

const (
	ERROR = iota
	EOF
	UNKNOWN
	WS // White Space
	STRING
	NUMBER
	WILDCARD
	OPERATOR
	
	BOOLOP // Boolean Operator

	COMMENT

	// ()
	OPENGRP
	CLOSEGRP

	SELECT
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
	initPosition := l.Cursor-1
	nextSymbol := l.next()
	
	for b(nextSymbol) {
		nextSymbol = l.next()
	}

	return l.Code[initPosition:l.Cursor-1]
}

func (l *Lex) skipWhitespace() {
	for unicode.IsSpace(rune(l.Code[l.Cursor])) {
		l.next()
	}
}


func (l *Lex) GetToken() Token {
	var token Token

	l.skipWhitespace()
	if int(l.Cursor) == len(l.Code) {
		token = Token{Data: "", Type: EOF}
	}

	data := l.next()

	switch data {
	// First analyze single rune tokens
	case '=' :
		token = Token{Data: "", Type: BOOLOP}
	case 0:
		token = Token{Data: "", Type: EOF}
	default:
		if isLetter(data){
			literal := l.scanFullSequence(isLetter)
			tokenType := LookupToken(literal) 
			token = Token{Data: literal, Type: tokenType}
			
		} else if (isDigit(data)) {
			literal := l.scanFullSequence(isDigit)
			var tokenType uint = NUMBER // token.Type
			token = Token{Data: literal, Type: tokenType}
		} else {
			token = Token{Data: "", Type: UNKNOWN}
		}
	}

	return token
}
