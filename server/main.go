package main

import (
	"net/http"
)

func main() {
	s := CreateNewServer()
	s.MountHandlers()
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })

	// fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", s.Router)

}
