package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	upstream 	string
	httpHost 	string
	httpsHost 	string
	path 		string
	key 		string
	crt 		string
}

/*
	New function configures the upstream based on the passed one
*/
func New(ups string) Server {
	s := Server{
		upstream: ups,
		httpHost: "0.0.0.0:8080",
		httpsHost: "0.0.0.0:4343",
		path: "/",
		key: "server.key",
		crt: "server.crt",
	}
	return s
}

func (s *Server)SetHTTPHost(host string) {
	s.httpHost = host
}

func (s *Server)SetHTTPSHost(host string) {
	s.httpsHost = host
}

func (s *Server)SetCertificate(key,crt string) {
	s.key = key
	s.crt = crt
}

func (s Server)createServer() (*http.Server,*http.Server) {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc(s.path, s.analyzer(s.proxy(s.upstream)))

	
	server := &http.Server{
		Addr: 		s.httpHost,
		Handler: 	serverMux,
	}

	serverTLS := &http.Server{
		Addr: 		s.httpsHost,
		Handler: 	serverMux,
	}

	return server,serverTLS
}

func (s Server)StartServer() {
	s1, s2 := s.createServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Listening HTTP")
		if err := s1.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Listener HTTP failed: %v\n", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Listening HTTPS")
		if err := s2.ListenAndServeTLS(s.crt, s.key); err != nil && err != http.ErrServerClosed {
			log.Printf("Listener HTTPS failed: %v\n", err)
		}
	}()

	<-sigChan
	log.Println("Shutdown signal received...")

	shutdownServer(ctx, s1)
	shutdownServer(ctx, s2)

	cancel()
	wg.Wait()
}

func shutdownServer(ctx context.Context, srv *http.Server) {
	gracefulCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(gracefulCtx); err != nil {
		log.Printf("Server shutdown failed: %v\n", err)
	} else {
		log.Println("Server shutdown OK")
	}
}
