package main

import (
	"os"
	"tutorial/server"

	log "github.com/sirupsen/logrus"
	_ "google.golang.org/grpc"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	logger := log.WithFields(log.Fields{})
	logger.Info("Example API")
}
func main() {
	server.NewHTTPServer()
}
