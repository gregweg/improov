package routes

import (
	"improov/handlers"
	"improov/middleware"
	"improov/storage"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	taskHandler := &handlers.TaskHandler{DB: &storage.DBImpl{}}

	// Public route: login
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware)

	api.HandleFunc("/api/categories", taskHandler.ListCategories).Methods("GET")
	api.HandleFunc("/api/tasks/suggest", taskHandler.SuggestTask).Methods("GET")
	api.HandleFunc("/api/tasks/complete", taskHandler.CompleteTask).Methods("POST")
	api.HandleFunc("/api/tasks/completed", taskHandler.GetCompletedTasks).Methods("GET")
}
