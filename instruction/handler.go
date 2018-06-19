package instruction

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	M Manager
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/instruction", h.getInstruction).Methods("GET")
}

func (h *Handler) getInstruction(w http.ResponseWriter, r *http.Request) {
	//
	w.WriteHeader(http.StatusNoContent)
}
