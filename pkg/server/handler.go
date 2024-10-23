package server

import (
	"io"
	"log"
	"net/http"
	"time"
)

/*
	HandlerFunc is a adapter. It allows us to pass a function that will receive
	the traffic on the second stage. Fist we gonna analyze the package with the
	parser, implementing for all those parameters that must be suitable for tests.

	If it could not spot irregularities for that request it must be passed to the proxy.


	TODO: change name to ServeHTTP and use as handler.
*/
func (s Server)analyzer(prx http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		attack := false
		queryParams := r.URL.Query()
		for _, values := range queryParams {
			for _, value := range values {
				if s.waf.Test(value) && !attack {attack = true}
			}
		}
		if attack {
			w.WriteHeader(403)
			w.Write([]byte("Malicious request detected!"))
			return
		}

		prx(w, r)
	}
}

/*
	Passthrough proxy. Mirror the interaction of the client with the upstream.
	The body must update in chunks.
*/
func (s Server)proxy(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpoint := target+r.RequestURI
		req, err := http.NewRequest(r.Method, endpoint, r.Body)
		if err != nil {
			// implement log pkg for all errors
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		log.Printf("edp=%s adr=%s uri=%s a=%s\n",endpoint, r.RemoteAddr, r.RequestURI, r.URL.Query().Get("a"))

		req.Header = r.Header // mirror headers
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		CompileXForwardHead(req.Header, r.RemoteAddr)
		
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		MirrorHeader(w.Header(), resp.Header)
		w.WriteHeader(resp.StatusCode)

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}
}