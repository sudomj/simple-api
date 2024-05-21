package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/x1bdev/simple-api/pkg/configuration/logger"
	"github.com/x1bdev/simple-api/pkg/controller"
)

func main() {

	logger := logger.NewLogger("INFO")
	logger.Configure()
	server := http.NewServeMux()

	healthController := controller.NewHealthController()
	server.HandleFunc("/health", healthController.GetHealth)

	port := os.Getenv("APP_PORT")
	slog.Info("Server running",
		"port", port)
	http.ListenAndServe(port, server)
}
