package controllers

import (
	"authorization-api/database"
	"authorization-api/services"
	"encoding/json"
	"net/http"
)

func UserDetails(writer http.ResponseWriter, request *http.Request) {
	userId := request.URL.Query().Get("userId")

	userService := &services.UserService{
		Connection: database.GetDatabaseConnection(),
		Context:    request.Context(),
	}

	userDetails, err := userService.GetUserDetails(userId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(userDetails)
}
