package controllers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	authGroup := router.PathPrefix("/auth").Subrouter()
	authGroup.HandleFunc("/token", AuthenticationEndpoint).Methods("POST")

	usersGroup := router.PathPrefix("/users").Subrouter()
	usersGroup.HandleFunc("/details", UserDetails).Methods("GET")

	healthGroup := router.PathPrefix("/health").Subrouter()
	healthGroup.HandleFunc("/ping", PingReply).Methods("GET")

	return router
}
