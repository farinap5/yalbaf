package server

import (
	"io"
	"net/http"
	"time"
)

/*
	HandlerFunc is a adapter. It allows us to pass a function that will receive
	the trafic on the second stage. Fist we gonna analyse the package with the
	parser, implementing for all those parameters that must be suitable for tests.

	If it could not spot irregularities for that request it must be passed to the proxy.


	TODO: change name to ServeHTTP and use as handler.
*/
func analyser(prx http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/* 
			Implement everything here
			
			w.WriteHeader(403)
			io.Copy(w, bloq page)
		*/
		prx(w, r)
	}
}

/*
	Passthrough proxy. Mirror the interation of the client with the upstream.
	The body must update in chunks.
*/
func proxy(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpoint := target+r.URL.RawPath
		req, err := http.NewRequest(r.Method, endpoint, r.Body)
		if err != nil {
			// implement log pkg for all errors
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		req.Header = r.Header // mirror headers
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
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