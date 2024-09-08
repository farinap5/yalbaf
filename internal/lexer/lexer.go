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
	Token	Token
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

func (l *Lex) GetToken() {

	if int(l.Cursor) == len(l.Code) {
		l.Token = Token{Data: "", Type: EOF}
	}

	data := l.next()

	for int(l.Cursor) < len(l.Code) {
		if isBlank(data) {
			l.Token = l.scan(isBlank, WS)
		} 
	}	
}
