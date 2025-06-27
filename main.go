package main

import (
	"authorization-api/controllers"
	"authorization-api/utils"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	slog.Info("Application started")
	utils.LoadEnv("")
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	err := http.ListenAndServe(":8080", controllers.Router())
	if err != http.ErrServerClosed {
		slog.Error("Failed to start server", "error", err)
		return
	}

	slog.Info("Application finished")
}
