package main

import (
	"fmt"

	"github.com/farinap5/yalbaf/internal/lexer"
)

func main() {
	exp := "SELECT/* um*cmt */xa/* um*cmt,y FROM x;"
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

	t = l.GetToken()
	fmt.Printf("Token: %s | Type: %d\n", t.Data, t.Type)
	
	t = l.GetToken()
	fmt.Printf("Token: %s | Type: %d\n", t.Data, t.Type)
}