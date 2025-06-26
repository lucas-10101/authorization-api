package main

import (
	"authorization-api/database"
	"authorization-api/services"
	"authorization-api/utils"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func main() {
	slog.Info("Application started")
	utils.LoadEnv()
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	authenticationService := &services.AuthenticationService{
		Connection: database.GetDatabaseConnection(),
		Context:    context.TODO(),
	}

	I, E := authenticationService.AuthenticateUser("LUCAS", "TESTE", "533b40c1-ec88-47a8-b12d-4619fbd72bc4")

	fmt.Println("I:", I, "E:", E)

	slog.Info("Application finished")
}
