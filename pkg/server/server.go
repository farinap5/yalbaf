package server

import (
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func createServer() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", analyser(proxy("http://google.com.br/")))

	
	server := &http.Server{
		Addr: 		"0.0.0.0:80",
		Handler: 	serverMux,
	}

	serverTLS := &http.Server{
		Addr: 		"0.0.0.0:443",
		Handler: 	serverMux,
	}
}

