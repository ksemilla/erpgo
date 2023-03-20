package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNewServer(db *mongo.Client) *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.DB = db
	return s
}

func LoadEnvVars() {
	log.Printf("Loading environment variables")
	godotenv.Load("../.env")
}
