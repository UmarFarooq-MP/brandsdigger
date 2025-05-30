package http

import (
	"brandsdigger/internal/interface/http/dto"
	"brandsdigger/internal/service"
	"encoding/json"
	"net/http"
)

type NamesHandler struct {
	nameService *service.NamesService
}

func NewNamesHandler(nameService *service.NamesService) *NamesHandler {
	return &NamesHandler{
		nameService: nameService,
	}
}

func (h *NamesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var namesRequest dto.NamesRequest

	if err := json.NewDecoder(r.Body).Decode(&namesRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	names, err := h.nameService.GetNames(namesRequest.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(names)
}
