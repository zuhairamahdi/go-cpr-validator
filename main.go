package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"bahrain-cpr/internal/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type validateRequest struct {
	CPR string `json:"cpr"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Post("/validate", handleValidateCPR)

	if err := http.ListenAndServe(":8080", r); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func handleValidateCPR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req validateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	cpr := strings.TrimSpace(req.CPR)
	if cpr == "" {
		writeError(w, http.StatusBadRequest, "CPR is required")
		return
	}

	ok, err := validator.ValidateCPR(cpr)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if !ok {
		writeError(w, http.StatusUnprocessableEntity, "CPR is invalid")
		return
	}

	_ = json.NewEncoder(w).Encode(validateResponse{Valid: true})
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(validateResponse{Valid: false, Error: message})
}
