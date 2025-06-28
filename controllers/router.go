package controllers

import (
	"authorization-api/controllers/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	authGroup := router.PathPrefix("/auth").Subrouter()
	authGroup.HandleFunc("/token", PostTokenForm).Methods(http.MethodPost, http.MethodOptions)

	usersGroup := router.PathPrefix("/users").Subrouter()
	usersGroup.Use(middlewares.BearerTokenMiddleware)
	usersGroup.HandleFunc("/details", GetUserDetails).Methods(http.MethodGet, http.MethodOptions)

	healthGroup := router.PathPrefix("/health").Subrouter()
	healthGroup.Use(middlewares.BearerTokenMiddleware)
	healthGroup.HandleFunc("/ping", GetPingReply).Methods(http.MethodGet, http.MethodOptions)

	return router
}

func SetupMiddleware(router *mux.Router) *mux.Router {

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
