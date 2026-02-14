package handler

import (
	"cloud-app-api/internal/service"
	"encoding/json"
	"net/http"
)

type ItemHandler struct {
	svc *service.ItemService
}

func NewItemHandler(s *service.ItemService) *ItemHandler {
	return &ItemHandler{svc: s}
}

func (h *ItemHandler) Items(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "missing name", 400)
			return
		}

		if err := h.svc.Add(name); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:

		items, err := h.svc.List()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(items)

	default:
		http.Error(w, "method not allowed", 405)
	}
}
