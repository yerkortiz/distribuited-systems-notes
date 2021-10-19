package server

import (
	"milestone3/handlers"
	"net/http"
)

func RunServer() {
	r := handlers.NewRouter()
	http.ListenAndServe(":8080", r)
}
