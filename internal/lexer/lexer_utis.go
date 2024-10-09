package lexer

import (
	"strings"
	"unicode"
)

var SymTable = map[string]uint{
	"SELECT": 	SELECT,
	"FROM": 	FROM,
	"WHERE": 	WHERE,
	"UNION":	UNION,
	"LIMIT":	LIMIT,
	"OR":		BOOLOP,
	"AND":		BOOLOP,
	"AS":		AS,
	"GROUP":	GROUP,
	"BY":		BY,
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
	meaning := SymTable[strings.ToUpper(token)]
	if meaning == 0 {
		return IDENTIFIER
	}
	return meaning
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