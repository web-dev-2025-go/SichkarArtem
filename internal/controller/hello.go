package controller

import (
	"go-http/internal/service"
	"net/http"
)

type Controller struct {
	mux     *http.ServeMux
	service service.Service
}

func NewController() {
}
