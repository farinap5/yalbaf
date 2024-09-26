package lexer_test

import (
	"testing"
	"github.com/farinap5/yalbaf/internal/lexer"

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
	l := lexer.Init("SELECT xa,y FROM any;")
	
	token := l.GetToken()
	t.Logf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 15 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	t.Logf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 4 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	t.Logf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 12 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	t.Logf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 4 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	t.Logf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 16 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	t.Logf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 4 {
		t.Error("Not the expected token type")
	}

	token = l.GetToken()
	t.Logf("Token: %s | Type: %d\n", token.Data, token.Type)
	if token.Type != 9 {
		t.Error("Not the expected token type")
	}
}