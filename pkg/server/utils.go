package server

import (
	"math/rand"
	"net/http"
	"strings"
)

func CompileXForwardHead(head http.Header, clientIP string) {
	if prior, ok := head["X-Forwarded-For"]; ok {
		clientIP = strings.Join(prior, ", ") + ", " + clientIP
	}
	head.Set("X-Forwarded-For", clientIP)
}

// Generate unique request ID. It is used mostly for
// debugging purpose.
func GenRequestID(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Add request id in request for the back end
func AddRequestID(head *http.Header) {
	head.Set("X-XYZ-ID",GenRequestID(16))
}

func MirrorHeader(dst, src http.Header) {
	for i, v := range src {
		for _, vlue := range v {
			dst.Add(i,vlue)
		}
	}
}
