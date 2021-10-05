package router

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloHandler).Methods("GET")
	router.HandleFunc("/compress", Compress).Methods("GET")
	return router
}
