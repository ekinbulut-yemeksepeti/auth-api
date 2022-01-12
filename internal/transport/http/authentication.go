package http

import (
	"encoding/json"
	"net/http"

	"github.com/ekinbulut-yemeksepeti/auth-api/internal/authentication"
)

func (h *Handler) CreateJWTToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var authentication authentication.Authentication
	if err := json.NewDecoder(r.Body).Decode(&authentication); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
	}

	token, err := h.Service.CreateJWTToken(&authentication)
	if err != nil {
		sendErrorResponse(w, "Failed to create JSON", err)
	}

	if err := json.NewEncoder(w).Encode(token); err != nil {
		panic(err)
	}
}
