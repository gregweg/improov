package main

import (
	"log"
	"net/http"

	"improov/routes"
	"improov/storage"

	"github.com/gorilla/mux"
)

func main() {
	storage.Init()
	storage.SeedTasks()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
