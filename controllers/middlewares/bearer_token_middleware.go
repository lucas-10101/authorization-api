package middlewares

import (
	"authorization-api/database"
	"authorization-api/services"
	"authorization-api/utils"
	"context"
	"net/http"
	"strings"
)

func BearerTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() { // token input error handling
			if r := recover(); r != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte("Internal Server Error"))
			}
		}()

		authenticationService := services.AuthenticationService{
			Connection: database.GetDatabaseConnection(),
			Context:    request.Context(),
		}

		token := request.Header.Get(utils.HeaderAuthorization)

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		token = strings.TrimSpace(token[7:])

		tokenData, err := authenticationService.ValidateToken(token)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		request = request.WithContext(
			context.WithValue(request.Context(),
				utils.RequestContextPrincipalKey,
				tokenData,
			),
		)

		next.ServeHTTP(writer, request)
	})
}
