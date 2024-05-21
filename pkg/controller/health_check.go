package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type healthController struct {
}

func NewHealthController() *healthController {
	return &healthController{}
}

func (hc *healthController) GetHealth(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"code":      0,
		"status":    "UP",
		"timestamp": time.Now(),
	}

	responseBytes, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slog.Info("health check",
		"data", response)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
