package web

import "net/http"

// Handler is defined as any function that accepts a request
// and returns a response
type Handler func(w http.ResponseWriter, r *http.Request)

// Service has a method handler, a rest api path
type Service interface {
	GetAPI() string
	GetHandler() http.HandlerFunc
	GetMethod() string
}
