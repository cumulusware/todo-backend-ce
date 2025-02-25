package todos

import (
	"net/http"

	"github.com/cumulusware/todo-backend-ce/internal/rest"
)

// DescribeAll handles the OPTIONS method for the /todos/ endpoint.
func DescribeAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.RespondWithOptions(w, "GET,POST,DELETE,OPTIONS")
	}
}

// Describe handles the OPTIONS method for the /todos/{key}/ endpoint.
func Describe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.RespondWithOptions(w, "GET,PATCH,DELETE,OPTIONS")
	}
}

// ReadAll handles the GET method to list all todos.
func ReadAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rest.RespondWithJSON(w, http.StatusOK, todos)
	}
}

// Create handles the POST method to create a new todo.
func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todo := struct {
			Title string `json:"title"`
		}{
			Title: "a todo",
		}
		rest.RespondWithJSON(w, http.StatusOK, todo)
	}
}

// DeleteAll handles the DELETE method to delete all todos.
func DeleteAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos = []Todo{}
		rest.RespondWithJSON(w, http.StatusNoContent, todos)
	}
}
