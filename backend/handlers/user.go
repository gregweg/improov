package handlers

import (
	"encoding/json"
	"improov/storage"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Future: return full user profile
	w.Write([]byte("User API not implemented yet"))
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User") // set in your JWT middleware
	if userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	db := &storage.DBImpl{}
	user, err := db.GetUserByID(userID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
