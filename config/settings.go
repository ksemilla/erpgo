package config

import (
	"log"
	"os"
)

type Settings struct {
	Port        string
	MongoDB_Uri string
}

func GetSettings() *Settings {
	s := &Settings{}

	port := os.Getenv("PORT")
	mongodb_uri := os.Getenv("MONGODB_URI")

	if port == "" {
		s.Port = "3000"
	} else {
		s.Port = port
	}
	s.MongoDB_Uri = mongodb_uri

	if mongodb_uri == "" {
		log.Fatal("MONGODB_URI environment variable missing")
	}

	return s
}
