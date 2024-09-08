package lexer

import "unicode"

func isLetter(t rune) bool {
	return unicode.IsLetter(t)
}

func isBlank(char rune) bool {
	switch char {
	case '\n':
		return true
	case '\t':
		return true
	case ' ':
		return true
	default:
		return false
	}
}

func (l *Lex) scan(b func(rune) bool, Type uint) Token {
	initPosition := l.Cursor
	nextSymbol := l.next()
	
	for b(nextSymbol) {
		nextSymbol = l.next()
	}

	return Token{Type: WS, Data: l.Code[initPosition:l.Cursor]}
}