package handlers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type HelloWorldHandler struct {
	DB *mongo.Client
}

// HelloWorld api Handler
func (h *HelloWorldHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	// h.DB
	w.Write([]byte("Hello World1!"))
}
