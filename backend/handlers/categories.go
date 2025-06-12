package handlers

import (
	"encoding/json"
	"improov/models"
	"net/http"
)

func (h *TaskHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.DefaultCategories)
}
