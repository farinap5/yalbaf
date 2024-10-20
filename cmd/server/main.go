package main

import "github.com/farinap5/yalbaf/pkg/server"

func main() {
	s := server.New("https://farinap5.com")
	s.StartServer()
}
