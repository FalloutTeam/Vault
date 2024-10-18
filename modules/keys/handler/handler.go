package handler

import (
	"encoding/json"
	"net/http"
	"vault/modules/keys/service"

	"github.com/go-chi/chi"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) AddSecret(w http.ResponseWriter, r *http.Request) {

	path := chi.URLParam(r, "path")

	if path == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// check verification

	var inputData map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	outputData, err := json.Marshal(inputData)

	if err != nil {
		http.Error(w, "Failed to encode response JSON", http.StatusInternalServerError)
		return
	}

	w.Write(outputData)

}

func (h *Handler) DeleteSecret(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "path")

	if path == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// check root

	// delete by path

}

func (h *Handler) PatchSecret(w http.ResponseWriter, r *http.Request) {

	path := chi.URLParam(r, "path")

	if path == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// check root

	// take Secret

	// decrypt

}

func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request) {

	path := chi.URLParam(r, "path")

	if path == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	// check root

	// take Secret

	// decrypt

}
