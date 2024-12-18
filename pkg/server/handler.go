package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

/*
	HandlerFunc is an adapter. It allows us to pass a function that will receive
	the traffic on the second stage. Fist we gonna analyze the package with the
	parser, implementing for all those parameters that must be suitable for tests.

	If it could not spot irregularities for that request it must be passed to the proxy.


	TODO: change name to ServeHTTP and use as handler.
*/
func (s Server)analyzer(prx http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		attack := false
		tokenFound := 0
		queryParams := r.URL.Query()
		for _, values := range queryParams {
			for _, value := range values {

				if s.vector == "str" {
					tokenFound, attack = s.waf.TestStr(value) // && !attack {attack = true}
				} else if s.vector == "int" {
					tokenFound, attack = s.waf.TestInt(value) // && !attack {attack = true}
				} else {
					tokenFound, attack = s.waf.TestStr(value) // && !attack {attack = true}
				}

			}
		}
		if attack {
			rid := GenRequestID(24)
			
			t := time.Now()
			w.Header().Set("X-XYZ-ID",rid)
			w.WriteHeader(403)

			w.Write([]byte(fmt.Sprintf(`Blocked due to malicious request detected!
Tokens found: %d
Your IP: %s
Request ID: %s
Time: %s
`, tokenFound, r.RemoteAddr, rid, t.Format("01-02-2006 15:04:05"))))
			endpoint := s.upstream+r.RequestURI
			log.Printf("edp=%s adr=%s uri=%s err=\"Exploitation attempt\" tks=%d rid=%s\n",
				endpoint, r.RemoteAddr, r.RequestURI,tokenFound ,rid)
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
		rid := GenRequestID(24)
		endpoint := target+r.RequestURI
		req, err := http.NewRequest(r.Method, endpoint, r.Body)
		if err != nil {
			// implement log pkg for all errors
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		log.Printf("edp=%s adr=%s uri=%s",endpoint, r.RemoteAddr, r.RequestURI)

		req.Header = r.Header // mirror headers
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		CompileXForwardHead(req.Header, r.RemoteAddr)
		req.Header.Add("X-XYZ-ID", rid)
		
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		MirrorHeader(w.Header(), resp.Header)
		w.Header().Set("X-XYZ-ID", rid)
		w.WriteHeader(resp.StatusCode)

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}
}