package main

import (
	"erpgo/config"
	"log"
	"net/http"
)

func main() {
	LoadEnvVars()

	settings := config.GetSettings()

	log.Println("Setting up server")
	s := CreateNewServer()
	s.Initialize()

	log.Printf("Listening on port %s...", settings.Port)
	log.Fatal(http.ListenAndServe(":"+settings.Port, s.Router))

}
