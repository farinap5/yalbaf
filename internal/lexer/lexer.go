package lexer

const (
	ERROR = iota
	EOF
	UNKNOWN
	WS // White Space
	IDENTIFIER
	STRING
	NUMBER
	WILDCARD
	OPERATOR
	DOTCOMMA
	
	BOOLOP // Boolean Operator AND OR

	COMMENT
	COMMA
	PLUS
	DOT

	// ()
	OPENGRP
	CLOSEGRP
	EQUAL
	GT
	LT
	GTE
	LTE
	NE

	SELECT
	FROM
	WHERE
	UNION
	LIMIT
	AS
	GROUP
	BY
)

type Token struct {
	Type uint
	Data string
}

type Lex struct {
	Code 	string
	Cursor 	uint
	CChar	byte
	Ptr   	uint
	//Token	Token
}


/*
	Init will create a new lexer for a new sentence
*/
func Init(Data string) *Lex {
	lex := new(Lex)
	lex.Code = Data
	lex.Cursor = 0
	lex.Ptr = 0
	lex.next()
	return lex
}

func (l *Lex) next() {
	if int(l.Cursor) >= len(l.Code) {
		l.CChar = 0
	} else {
		l.CChar = l.Code[l.Cursor]
	}
	l.Ptr = l.Cursor
	l.Cursor++
}

func (l *Lex) scanFullSequence(b func(rune) bool) string {
	initPosition := l.Ptr
	for b(rune(l.CChar)) {
		l.next()
	}

	return l.Code[initPosition:l.Ptr]
}

func (l *Lex) skipWhitespace() {
	for isBlank(rune(l.CChar)) {
		l.next()
	}
}


func (l *Lex) GetToken() Token {
	var token Token

	if int(l.Cursor) == len(l.Code) {
		token = Token{Data: "", Type: EOF}
	}

		
	l.skipWhitespace()

	switch l.CChar {
	// First analyze single rune tokens
	case ',' :
		token = Token{Data: ",", Type: COMMA}
	case ';' :
		token = Token{Data: ";", Type: DOTCOMMA}
	case '(' :
		token = Token{Data: "(", Type: OPENGRP}
	case ')' :
		token = Token{Data: ")", Type: CLOSEGRP}
	case '"','\'','`' :
		token = Token{Data: ")", Type: STRING}
	case '+' :
		token = Token{Data: "+", Type: PLUS}
	case '.' :
		token = Token{Data: ".", Type: DOT}
	case '>' :
		token = Token{Data: ">", Type: GT}
	case '<' :
		token = Token{Data: "<", Type: LT}
	case '=' :
		token = Token{Data: "=", Type: EQUAL}
	case 0:
		token = Token{Data: "", Type: EOF}
	default:
		if isLetter(rune(l.CChar)){
			literal := l.scanFullSequence(isLetter)
			tokenType := LookupToken(literal) 
			token = Token{Data: literal, Type: tokenType}
		} else if (isDigit(rune(l.CChar))) {
			literal := l.scanFullSequence(isDigit)
			var tokenType uint = NUMBER // token.Type
			token = Token{Data: literal, Type: tokenType}
		} else {
			token = Token{Data: "", Type: UNKNOWN}
		}
		return token
	}

	l.next()
	return token
}
