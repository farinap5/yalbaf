package main

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
)

func main() {
	exp := "SELECT x"
	l := lexer.Init(exp)
	t := l.GetToken()
	fmt.Printf("Token: %s | Type: %d", t.Data, t.Type)
}