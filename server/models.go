package main

import (
	"erpgo/handlers"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Router *chi.Mux
	DB     *mongo.Client
}

func (s *Server) MountHandlers() {
	// Mount all handlers here
	hellohandler := &handlers.HelloWorldHandler{DB: s.DB}
	s.Router.Get("/", hellohandler.HelloWorld)
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
