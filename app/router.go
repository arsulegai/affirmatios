package app

import (
	"affirmatios/hospital/web"

	"github.com/go-chi/chi"
)

// CustomRouter is to wrap and add utility methods
type CustomRouter struct {
	router *chi.Mux
}

// GetRouter returns the internal router
func (c *CustomRouter) GetRouter() *chi.Mux {
	return c.router
}

func (c *CustomRouter) setupSessionStore() {
	return
}

func (c *CustomRouter) setupRoutes(services []web.Service) {
	for _, service := range services {
		c.router.Get(service.GetAPI(), service.GetHandler())
	}
	return
}
