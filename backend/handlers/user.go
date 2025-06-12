package handlers

import (
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Future: return full user profile
	w.Write([]byte("User API not implemented yet"))
}
