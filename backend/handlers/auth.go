package handlers

import (
	"encoding/json"
	"improov/auth"
	"improov/models"
	"improov/storage"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB storage.DBInterface
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate fields
	if req.Username == "" || req.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Lookup user
	user, err := h.DB.GetUserByID(req.Username)
	if err != nil || user == nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return token as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" || req.Name == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	existing, _ := h.DB.GetUserByID(req.Username)
	if existing != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	user := &models.User{
		ID:       req.Username,
		Name:     req.Name,
		Password: string(hashed),
	}
	if err := h.DB.SaveUser(user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	token, _ := auth.GenerateJWT(req.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
