package server

import (
	"net/http"
	"tutorial/router"

	log "github.com/sirupsen/logrus"
)

func NewHTTPServer() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
