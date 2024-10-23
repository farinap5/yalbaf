package main

import "github.com/farinap5/yalbaf/pkg/server"

func main() {
	s := server.New("http://0.0.0.0:5555")
	s.StartServer()
}
