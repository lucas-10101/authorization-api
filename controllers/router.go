package controllers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/authenticate", AuthenticationEndpoint).Methods("POST")

	return router
}
