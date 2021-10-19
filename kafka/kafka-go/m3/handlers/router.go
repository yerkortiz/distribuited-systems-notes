package handlers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", HealthCheck).Methods("GET")
	router.HandleFunc("/order", ReceiveOrder).Methods("POST")
	return router
}
