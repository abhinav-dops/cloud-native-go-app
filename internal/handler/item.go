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
		var data struct {
			Name string `json:"name"`
		}
		json.NewDecoder(r.Body).Decode(&data)
		item := h.svc.CreateItem(data.Name)
		json.NewEncoder(w).Encode(item)

	case http.MethodGet:
		items := h.svc.ListItems()
		json.NewEncoder(w).Encode(items)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
