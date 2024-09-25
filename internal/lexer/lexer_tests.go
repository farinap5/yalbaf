package lexer

import (
	"testing"
	"log"
)

/*
	Token: SELECT | Type: 15
	Token: xa     | Type: 4
	Token: ,      | Type: 12
	Token: y      | Type: 4
	Token: FROM   | Type: 16
	Token: any    | Type: 4
	Token: ;      | Type: 9
*/
func TestLexerBasic(t *testing.T) {
	l := Init("SELECT xa,y FROM any;")
	
	token := l.GetToken()
	log.Printf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 15 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	log.Printf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 10 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	log.Printf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 10 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	log.Printf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 10 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	log.Printf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 10 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	log.Printf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 10 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	log.Printf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 10 {
		t.Error("Not the expected token type")
	}
}