package main

import "github.com/go-chi/chi/v5"

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}
