package http

import (
	"brandsdigger/internal/domain/auth"
	"brandsdigger/internal/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	namesService *service.NamesService
}

func NewAuthHandler(nameService *service.NamesService) *AuthHandler {
	return &AuthHandler{namesService: nameService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest auth.Login

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("names")
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var signUpRequest auth.SignUp

	if err := json.NewDecoder(r.Body).Decode(&signUpRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("names")
}
