package config

import (
	"os"
)

type Settings struct {
	Port string
}

func GetSettings() *Settings {
	s := &Settings{}
	port := os.Getenv("PORT")
	if port == "" {
		s.Port = "3000"
	} else {
		s.Port = port
	}

	return s
}
