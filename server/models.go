package main

import (
	"erpgo/handlers"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func (s *Server) MountHandlers() {
	// Mount all handlers here
	s.Router.Get("/", handlers.HelloWorld)
}

func (s *Server) MountMiddlewares() {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)
}

func (s *Server) Initialize() {
	log.Println("Initializing server")
	s.MountMiddlewares()
	s.MountHandlers()
}
