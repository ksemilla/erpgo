package main

import (
	"erpgo/config"
	"erpgo/database"
	"log"
	"net/http"
)

func main() {
	LoadEnvVars()

	settings := config.GetSettings()

	db := database.CreateDBClient()

	s := CreateNewServer(db)
	s.Initialize()

	log.Printf("Listening on port %s...", settings.Port)
	log.Fatal(http.ListenAndServe(":"+settings.Port, s.Router))

}
