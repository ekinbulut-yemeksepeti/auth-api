package http

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/ekinbulut-yemeksepeti/auth-api/internal/authentication"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Service *authentication.Service
}

// Response objecgi
type Response struct {
	Message string
	Error   string
}


func NewHandler(service *authentication.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	log.Info("Setting up routes...")

	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/token", h.CreateJWTToken).Methods("POST")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
			panic(err)
		}
	})
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
