package controllers

import "net/http"

func AuthenticationEndpoint(writter http.ResponseWriter, request *http.Request) {

	//var username, password, tenantId string

	writter.WriteHeader(http.StatusUnauthorized)
	writter.Write([]byte("Unauthorized"))
}
