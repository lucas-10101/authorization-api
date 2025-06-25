package main

import (
	"authorization-api/utils"
	"log/slog"
)

func main() {
	slog.Info("Application started")

	utils.LoadEnv()
	slog.Info("Application finished")
}
