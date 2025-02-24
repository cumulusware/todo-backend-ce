package main

import (
	"net/http"

	"github.com/cumulusware/todo-backend-ce/internal/rest/todos"
	"github.com/rs/cors"
)

func createRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/api/todos/", http.StripPrefix("/api/todos", todos.AddRoutes()))
	return router
}

func setupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "OPTIONS", "POST", "DELETE", "PUT", "PATCH"},
		AllowedHeaders:   []string{"accept", "content-type"},
		AllowCredentials: true,
		Debug:            true,
	})
}
