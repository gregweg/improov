package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"improov/models"
	"improov/storage"
)

type TaskHandler struct {
	DB storage.DBInterface
}

func (h *TaskHandler) SuggestTask(w http.ResponseWriter, r *http.Request) {
	category := strings.ToLower(r.URL.Query().Get("category"))
	userID := r.URL.Query().Get("userId")

	if category == "" || userID == "" {
		http.Error(w, "Missing 'category' or 'userId'", http.StatusBadRequest)
		return
	}

	tasks, err := h.DB.GetTasksByCategory(category)
	if err != nil || len(tasks) == 0 {
		http.Error(w, "No tasks available", http.StatusNotFound)
		return
	}

	rand.Seed(time.Now().UnixNano())
	chosen := tasks[rand.Intn(len(tasks))]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chosen)

}

func (h *TaskHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID   string `json:"user_id"`
		TaskID   string `json:"task_id"`
		Category string `json:"category"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if payload.UserID == "" || payload.TaskID == "" || payload.Category == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	user, err := h.DB.GetOrCreateUser(payload.UserID)
	if err != nil {
		http.Error(w, "Failed to get or create user", http.StatusInternalServerError)
		return
	}

	// Update stats
	switch strings.ToLower(payload.Category) {
	case "fitness":
		user.Stats.Fitness += 10
	case "learning":
		user.Stats.Learning += 10
	case "mindfulness":
		user.Stats.Mindfulness += 10
	}

	//user.Points += 10 //

	// Save user and task completion
	if err := h.DB.SaveUser(user); err != nil {
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	if err := h.DB.CreateCompletedTask(payload.UserID, payload.TaskID); err != nil {
		http.Error(w, "Failed to create task completion", http.StatusInternalServerError)
		return
	}

	user.Stats = models.UserStats{
		Fitness:     user.Stats.Fitness,
		Learning:    user.Stats.Learning,
		Mindfulness: user.Stats.Mindfulness,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *TaskHandler) GetCompletedTasks(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	completed, err := h.DB.GetCompletedTasks(userID)
	if err != nil {
		http.Error(w, "Failed to fetch completed tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}
