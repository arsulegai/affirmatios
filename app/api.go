package app

import (
	"affirmatios/university/internal/management"
	"affirmatios/university/internal/university"
	"affirmatios/university/internal/user"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// Run will be the entry point function for the REST APIs
func Run(config *Config) error {
	universityHandler := university.University{}
	connectionsHandler := management.Connections{}
	userHandler := user.Management{}

	r := getRouter()
	r.setupSessionStore()
	// add application services
	r.setupRoutes(connectionsHandler.GetServices())
	// add application services
	r.setupRoutes(universityHandler.GetServices())
	// add user management services
	r.setupRoutes(userHandler.GetServices())
	// add configuration services
	r.setupRoutes(config.GetServices())
	address := fmt.Sprintf("%s:%s", config.GetHost(), config.GetPort())
	log.Printf("Started the service at %s\n", address)
	return http.ListenAndServe(address, r.GetRouter())
}

func getRouter() CustomRouter {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	customRouter := CustomRouter{
		r,
	}
	return customRouter
}
