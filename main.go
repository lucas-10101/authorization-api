package main

import (
	"authorization-api/utils"
	"log/slog"
	"time"
)

func main() {
	slog.Info("Application started")
	utils.LoadEnv("")
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	slog.Info("Application finished")
}
