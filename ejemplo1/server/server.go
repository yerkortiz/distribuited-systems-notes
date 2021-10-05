package server

import (
	"example/router"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func NewHTTPServer() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
