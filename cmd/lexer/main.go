package main

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
)

func main() {
	exp := "SELECT xa,y"
	l := lexer.Init(exp)
	t := l.GetToken()
	fmt.Printf("Token: %s | Type: %d\n", t.Data, t.Type)

	t = l.GetToken()
	fmt.Printf("Token: %s | Type: %d\n", t.Data, t.Type)

	t = l.GetToken()
	fmt.Printf("Token: %s | Type: %d\n", t.Data, t.Type)

	t = l.GetToken()
	fmt.Printf("Token: %s | Type: %d\n", t.Data, t.Type)

	t = l.GetToken()
	fmt.Printf("Token: %s | Type: %d\n", t.Data, t.Type)
	
}