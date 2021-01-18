package app

import (
	"affirmatios/university/internal/user"
	"affirmatios/university/web"
	"log"
	"net/http"

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
	c.router.Use(sessionMiddleware)
}

func sessionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoggedInSession := "user-logged-in"
		_, err := user.GetStore().Get(r, userLoggedInSession)
		if err != nil {
			log.Println(err)
			http.SetCookie(w, &http.Cookie{Name: userLoggedInSession, MaxAge: -1, Path: "/"})
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (c *CustomRouter) setupRoutes(services []web.Service) {
	for _, service := range services {
		if service.GetMethod() == http.MethodGet {
			c.router.Get(service.GetAPI(), service.GetHandler())
		}
		if service.GetMethod() == http.MethodPost {
			c.router.Post(service.GetAPI(), service.GetHandler())
		}
	}
	return
}
