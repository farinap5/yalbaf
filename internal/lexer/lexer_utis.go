package lexer

import (
	"strings"
	"unicode"
)

var Batata = map[string]uint{
	"SELECT": SELECT,
}

func isLetter(t rune) bool {
	return unicode.IsLetter(t)
}

func isLetterNumb(t rune) bool {
	return unicode.IsLetter(t) || unicode.IsNumber(t)
}

func isDigit(t rune) bool {
	return unicode.IsDigit(t)
}

func LookupToken(token string) uint {
	return Batata[strings.ToUpper(token)]
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