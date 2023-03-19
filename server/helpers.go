package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func LoadEnvVars() {
	log.Printf("Loading environment variables")
	godotenv.Load("../.env")
}
