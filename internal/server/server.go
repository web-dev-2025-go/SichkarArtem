package server

import (
	"context"
	"net/http"
)

type Http struct {
	server *http.Server
}

func NewHttp(addr string) *Http {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return &Http{
		server: server,
	}
}

func (h *Http) Start() error {
	return h.server.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
