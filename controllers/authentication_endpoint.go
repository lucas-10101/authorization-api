package controllers

import (
	"authorization-api/database"
	"authorization-api/services"
	"encoding/json"
	"net/http"
)

func AuthenticationEndpoint(writter http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	request.ParseMultipartForm(1024 * 32)

	username := request.PostFormValue("username")
	password := request.PostFormValue("password")
	tenantId := request.PostFormValue("tenantId")

	service := &services.AuthenticationService{
		Connection: database.GetDatabaseConnection(),
		Context:    request.Context(),
	}

	token, err := service.AuthenticateUser(username, password, tenantId)

	if err != nil {
		writter.WriteHeader(http.StatusUnauthorized)
		writter.Write([]byte("Unauthorized"))
		return
	}

	response := map[string]string{
		"access_token":  token,
		"id_token":      "",
		"refresh_token": "",
		"token_type":    "Bearer",
	}

	var bytesOut []byte
	if bytesOut, err = json.Marshal(response); err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		return
	}

	writter.Header().Set("Content-Type", "application/json")
	writter.Write([]byte(bytesOut))
}
