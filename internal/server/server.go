package server;

import "net/http"

type Http struct {
	server *http.Server
}

func NewHttp(addr string) *Http {
	mux := http.NewServeMux()

	handler := NewHandler()
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return &Http{
		server: 
	}
}
