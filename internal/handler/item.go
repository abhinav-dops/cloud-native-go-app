package handler

import (
	"cloud-app-api/internal/service"
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
			writeJSON(w, http.StatusBadRequest, APIResponse{
				Status: "error",
				Error:  "missing name",
			})
			return
		}

		if err := h.svc.Add(name); err != nil {
			writeJSON(w, 500, APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusCreated, APIResponse{
			Status: "success",
		})

	case http.MethodGet:
		items, err := h.svc.List()
		if err != nil {
			writeJSON(w, 500, APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		writeJSON(w, 200, APIResponse{
			Status: "success",
			Data:   items,
		})

	default:
		writeJSON(w, 405, APIResponse{
			Status: "error",
			Error:  "method not allowed",
		})
	}
}
