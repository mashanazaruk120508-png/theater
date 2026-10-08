package handler

import (
	"encoding/json"
	"net/http"
	"theater-backend/internal/models"
	"theater-backend/internal/service"
)

type PlayHandler struct {
	service *service.PlayService
}

func NewPlayHandler(service *service.PlayService) *PlayHandler {
	return &PlayHandler{service: service}
}

func (h *PlayHandler) GetAllPlays(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	plays, err := h.service.GetAllPlays(r.Context())
	if err != nil {
		http.Error(w, "Ошибка получения списка спектаклей", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plays)
}

func (h *PlayHandler) CreatePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var play models.Play
	if err := json.NewDecoder(r.Body).Decode(&play); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.CreatePlay(r.Context(), &play); err != nil {
		http.Error(w, "Ошибка создания спектакля", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(play)
}
