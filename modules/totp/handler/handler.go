package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"vault/modules/totp/handler/request"
	"vault/modules/totp/handler/response"
	"vault/modules/totp/models"
	"vault/modules/totp/service"
)

type Totp interface {
	CreateTotp(w http.ResponseWriter, r *http.Request)
	//TODO: Delete()
	//TODO: Read()
	//TODO: List()
}
type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (handler *Handler) CreateTotp(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	var body request.CreateTotpRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	barCode, url, err := handler.service.Totp.CreateTotp(name, models.Totp{
		Generate:    body.Generate,
		Exported:    body.Exported,
		KeySize:     body.KeySize,
		Issuer:      body.Issuer,
		AccountName: body.AccountName,
		Period:      body.Period,
		Algorithm:   body.Algorithm,
		Digits:      body.Digits,
		Skew:        body.Skew,
		QrSize:      body.QrSize,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	json.NewEncoder(w).Encode(response.CreateTotpResponse{Data: response.Data{
		Barcode: barCode,
		Url:     url,
	}})
}
